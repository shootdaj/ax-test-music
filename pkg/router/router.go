package router

import (
	"net/http"
	"strings"

	"github.com/shootdaj/ax-test-music/pkg/handlers"
)

// New creates a new HTTP handler with all routes configured
func New(h *handlers.Handler) http.Handler {
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/health", h.Health)
	mux.HandleFunc("/api/artists", h.Artists)
	mux.HandleFunc("/api/artists/", h.Artists)
	mux.HandleFunc("/api/albums", h.Albums)
	mux.HandleFunc("/api/albums/", h.Albums)
	mux.HandleFunc("/api/tracks", h.Tracks)
	mux.HandleFunc("/api/tracks/", h.Tracks)
	mux.HandleFunc("/api/playlists", h.Playlists)
	mux.HandleFunc("/api/playlists/", h.Playlists)
	mux.HandleFunc("/api/library", h.Library)
	mux.HandleFunc("/api/library/", h.Library)
	mux.HandleFunc("/api/search", h.SearchHandler)
	mux.HandleFunc("/api/browse", h.Browse)
	mux.HandleFunc("/api/browse/", h.Browse)
	mux.HandleFunc("/api/playback", h.Playback)
	mux.HandleFunc("/api/playback/", h.Playback)
	mux.HandleFunc("/api/stats", h.StatsHandler)

	// Frontend - serve the SPA for all non-API routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.NotFound(w, r)
			return
		}
		h.ServeFrontend(w, r)
	})

	// Add CORS middleware
	return corsMiddleware(mux)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
