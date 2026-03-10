package handler

import (
	"net/http"
	"sync"

	"github.com/shootdaj/ax-test-music/pkg/handlers"
	"github.com/shootdaj/ax-test-music/pkg/router"
	"github.com/shootdaj/ax-test-music/pkg/seed"
	"github.com/shootdaj/ax-test-music/pkg/store"
)

var (
	appHandler http.Handler
	once       sync.Once
)

func initApp() {
	s := store.New()
	seed.LoadSeedData(s)
	h := handlers.New(s)
	appHandler = router.New(h)
}

// Handler is the Vercel serverless function entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(initApp)
	appHandler.ServeHTTP(w, r)
}
