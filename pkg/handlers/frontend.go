package handlers

import (
	"embed"
	"net/http"
)

//go:embed frontend.html
var frontendFS embed.FS

// ServeFrontend serves the frontend SPA HTML
func (h *Handler) ServeFrontend(w http.ResponseWriter, r *http.Request) {
	data, err := frontendFS.ReadFile("frontend.html")
	if err != nil {
		http.Error(w, "frontend not found", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
