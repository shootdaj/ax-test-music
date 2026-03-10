package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/shootdaj/ax-test-music/pkg/models"
	"github.com/shootdaj/ax-test-music/pkg/store"
)

// Handler holds the store and provides HTTP handler methods
type Handler struct {
	Store *store.Store
}

// New creates a new Handler
func New(s *store.Store) *Handler {
	return &Handler{Store: s}
}

// writeJSON writes a JSON response
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// writeError writes a JSON error response
func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

// parseID extracts the ID from a URL path like /api/artists/artist-1
func parseID(path, prefix string) string {
	trimmed := strings.TrimPrefix(path, prefix)
	trimmed = strings.TrimPrefix(trimmed, "/")
	// Remove trailing slashes and any further path segments
	if idx := strings.Index(trimmed, "/"); idx >= 0 {
		trimmed = trimmed[:idx]
	}
	return trimmed
}

// --- Health ---

// Health handles GET /api/health
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// --- Artists ---

// Artists handles /api/artists and /api/artists/{id}
func (h *Handler) Artists(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// /api/artists/{id}/tracks
	if strings.Contains(path, "/tracks") {
		id := parseID(path, "/api/artists")
		id = strings.TrimSuffix(id, "/tracks")
		tracks := h.Store.GetTracksByArtist(id)
		if tracks == nil {
			tracks = []models.Track{}
		}
		writeJSON(w, http.StatusOK, tracks)
		return
	}

	// /api/artists/{id}/albums
	if strings.Contains(path, "/albums") {
		id := parseID(path, "/api/artists")
		id = strings.TrimSuffix(id, "/albums")
		albums := h.Store.GetAlbumsByArtist(id)
		if albums == nil {
			albums = []models.Album{}
		}
		writeJSON(w, http.StatusOK, albums)
		return
	}

	// /api/artists/{id}/related
	if strings.Contains(path, "/related") {
		id := parseID(path, "/api/artists")
		id = strings.TrimSuffix(id, "/related")
		artist, ok := h.Store.GetArtist(id)
		if !ok {
			writeError(w, http.StatusNotFound, "artist not found")
			return
		}
		// Find related artists by same genre
		var related []models.Artist
		for _, a := range h.Store.ListArtists() {
			if a.ID != artist.ID && a.Genre == artist.Genre {
				related = append(related, a)
			}
		}
		if related == nil {
			related = []models.Artist{}
		}
		writeJSON(w, http.StatusOK, related)
		return
	}

	id := parseID(path, "/api/artists")

	switch r.Method {
	case http.MethodGet:
		if id == "" {
			writeJSON(w, http.StatusOK, h.Store.ListArtists())
		} else {
			artist, ok := h.Store.GetArtist(id)
			if !ok {
				writeError(w, http.StatusNotFound, "artist not found")
				return
			}
			writeJSON(w, http.StatusOK, artist)
		}
	case http.MethodPost:
		var artist models.Artist
		if err := json.NewDecoder(r.Body).Decode(&artist); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		created := h.Store.CreateArtist(artist)
		writeJSON(w, http.StatusCreated, created)
	case http.MethodPut:
		if id == "" {
			writeError(w, http.StatusBadRequest, "id required")
			return
		}
		var artist models.Artist
		if err := json.NewDecoder(r.Body).Decode(&artist); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		artist.ID = id
		updated, ok := h.Store.UpdateArtist(artist)
		if !ok {
			writeError(w, http.StatusNotFound, "artist not found")
			return
		}
		writeJSON(w, http.StatusOK, updated)
	case http.MethodDelete:
		if id == "" {
			writeError(w, http.StatusBadRequest, "id required")
			return
		}
		if !h.Store.DeleteArtist(id) {
			writeError(w, http.StatusNotFound, "artist not found")
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"deleted": id})
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// --- Albums ---

// Albums handles /api/albums and /api/albums/{id}
func (h *Handler) Albums(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// /api/albums/{id}/tracks
	if strings.Contains(path, "/tracks") {
		id := parseID(path, "/api/albums")
		id = strings.TrimSuffix(id, "/tracks")
		tracks := h.Store.GetTracksByAlbum(id)
		if tracks == nil {
			tracks = []models.Track{}
		}
		writeJSON(w, http.StatusOK, tracks)
		return
	}

	id := parseID(path, "/api/albums")

	switch r.Method {
	case http.MethodGet:
		if id == "" {
			writeJSON(w, http.StatusOK, h.Store.ListAlbums())
		} else {
			album, ok := h.Store.GetAlbum(id)
			if !ok {
				writeError(w, http.StatusNotFound, "album not found")
				return
			}
			writeJSON(w, http.StatusOK, album)
		}
	case http.MethodPost:
		var album models.Album
		if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		created := h.Store.CreateAlbum(album)
		writeJSON(w, http.StatusCreated, created)
	case http.MethodPut:
		if id == "" {
			writeError(w, http.StatusBadRequest, "id required")
			return
		}
		var album models.Album
		if err := json.NewDecoder(r.Body).Decode(&album); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		album.ID = id
		updated, ok := h.Store.UpdateAlbum(album)
		if !ok {
			writeError(w, http.StatusNotFound, "album not found")
			return
		}
		writeJSON(w, http.StatusOK, updated)
	case http.MethodDelete:
		if id == "" {
			writeError(w, http.StatusBadRequest, "id required")
			return
		}
		if !h.Store.DeleteAlbum(id) {
			writeError(w, http.StatusNotFound, "album not found")
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"deleted": id})
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// --- Tracks ---

// Tracks handles /api/tracks and /api/tracks/{id}
func (h *Handler) Tracks(w http.ResponseWriter, r *http.Request) {
	id := parseID(r.URL.Path, "/api/tracks")

	switch r.Method {
	case http.MethodGet:
		if id == "" {
			writeJSON(w, http.StatusOK, h.Store.ListTracks())
		} else {
			track, ok := h.Store.GetTrack(id)
			if !ok {
				writeError(w, http.StatusNotFound, "track not found")
				return
			}
			writeJSON(w, http.StatusOK, track)
		}
	case http.MethodPost:
		var track models.Track
		if err := json.NewDecoder(r.Body).Decode(&track); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		created := h.Store.CreateTrack(track)
		writeJSON(w, http.StatusCreated, created)
	case http.MethodPut:
		if id == "" {
			writeError(w, http.StatusBadRequest, "id required")
			return
		}
		var track models.Track
		if err := json.NewDecoder(r.Body).Decode(&track); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		track.ID = id
		updated, ok := h.Store.UpdateTrack(track)
		if !ok {
			writeError(w, http.StatusNotFound, "track not found")
			return
		}
		writeJSON(w, http.StatusOK, updated)
	case http.MethodDelete:
		if id == "" {
			writeError(w, http.StatusBadRequest, "id required")
			return
		}
		if !h.Store.DeleteTrack(id) {
			writeError(w, http.StatusNotFound, "track not found")
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"deleted": id})
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// --- Playlists ---

// Playlists handles /api/playlists and /api/playlists/{id}
func (h *Handler) Playlists(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	// /api/playlists/{id}/tracks
	if strings.HasSuffix(path, "/tracks") || strings.Contains(path, "/tracks/") {
		h.PlaylistTracks(w, r)
		return
	}

	id := parseID(path, "/api/playlists")

	switch r.Method {
	case http.MethodGet:
		if id == "" {
			writeJSON(w, http.StatusOK, h.Store.ListPlaylists())
		} else {
			playlist, ok := h.Store.GetPlaylist(id)
			if !ok {
				writeError(w, http.StatusNotFound, "playlist not found")
				return
			}
			writeJSON(w, http.StatusOK, playlist)
		}
	case http.MethodPost:
		var playlist models.Playlist
		if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		created := h.Store.CreatePlaylist(playlist)
		writeJSON(w, http.StatusCreated, created)
	case http.MethodPut:
		if id == "" {
			writeError(w, http.StatusBadRequest, "id required")
			return
		}
		var playlist models.Playlist
		if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		playlist.ID = id
		updated, ok := h.Store.UpdatePlaylist(playlist)
		if !ok {
			writeError(w, http.StatusNotFound, "playlist not found")
			return
		}
		writeJSON(w, http.StatusOK, updated)
	case http.MethodDelete:
		if id == "" {
			writeError(w, http.StatusBadRequest, "id required")
			return
		}
		if !h.Store.DeletePlaylist(id) {
			writeError(w, http.StatusNotFound, "playlist not found")
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"deleted": id})
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// PlaylistTracks handles add/remove/reorder tracks in playlists
func (h *Handler) PlaylistTracks(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	// Extract playlist ID: /api/playlists/{id}/tracks
	parts := strings.Split(strings.TrimPrefix(path, "/api/playlists/"), "/")
	if len(parts) < 2 {
		writeError(w, http.StatusBadRequest, "invalid path")
		return
	}
	playlistID := parts[0]

	switch r.Method {
	case http.MethodGet:
		playlist, ok := h.Store.GetPlaylist(playlistID)
		if !ok {
			writeError(w, http.StatusNotFound, "playlist not found")
			return
		}
		// Return full track objects
		var tracks []models.Track
		for _, tid := range playlist.TrackIDs {
			if t, ok := h.Store.GetTrack(tid); ok {
				tracks = append(tracks, t)
			}
		}
		if tracks == nil {
			tracks = []models.Track{}
		}
		writeJSON(w, http.StatusOK, tracks)
	case http.MethodPost:
		var req struct {
			TrackID string `json:"track_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		if !h.Store.AddTrackToPlaylist(playlistID, req.TrackID) {
			writeError(w, http.StatusNotFound, "playlist or track not found")
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"added": req.TrackID})
	case http.MethodDelete:
		indexStr := r.URL.Query().Get("index")
		if indexStr == "" {
			writeError(w, http.StatusBadRequest, "index query param required")
			return
		}
		index, err := strconv.Atoi(indexStr)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid index")
			return
		}
		if !h.Store.RemoveTrackFromPlaylist(playlistID, index) {
			writeError(w, http.StatusNotFound, "playlist not found or invalid index")
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"removed": "true"})
	case http.MethodPut:
		var req struct {
			From int `json:"from"`
			To   int `json:"to"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		if !h.Store.ReorderPlaylistTrack(playlistID, req.From, req.To) {
			writeError(w, http.StatusBadRequest, "invalid reorder")
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"reordered": "true"})
	default:
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
	}
}

// --- Library ---

// Library handles /api/library endpoints
func (h *Handler) Library(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case strings.HasSuffix(path, "/like"):
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		var req struct {
			TrackID string `json:"track_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		if !h.Store.ToggleLike(req.TrackID) {
			writeError(w, http.StatusNotFound, "track not found")
			return
		}
		liked := h.Store.IsLiked(req.TrackID)
		writeJSON(w, http.StatusOK, map[string]interface{}{"track_id": req.TrackID, "liked": liked})

	case strings.HasSuffix(path, "/liked"):
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		trackIDs := h.Store.GetLikedTracks()
		var tracks []models.Track
		for _, tid := range trackIDs {
			if t, ok := h.Store.GetTrack(tid); ok {
				tracks = append(tracks, t)
			}
		}
		if tracks == nil {
			tracks = []models.Track{}
		}
		writeJSON(w, http.StatusOK, tracks)

	case strings.HasSuffix(path, "/recent"):
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		trackIDs := h.Store.GetRecentlyPlayed()
		var tracks []models.Track
		for _, tid := range trackIDs {
			if t, ok := h.Store.GetTrack(tid); ok {
				tracks = append(tracks, t)
			}
		}
		if tracks == nil {
			tracks = []models.Track{}
		}
		writeJSON(w, http.StatusOK, tracks)

	default:
		// /api/library - return full library summary
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"liked_count":  len(h.Store.GetLikedTracks()),
			"recent_count": len(h.Store.GetRecentlyPlayed()),
		})
	}
}

// --- Search ---

// SearchHandler handles /api/search?q=query
func (h *Handler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	query := r.URL.Query().Get("q")
	if query == "" {
		writeJSON(w, http.StatusOK, models.SearchResults{
			Artists: []models.Artist{},
			Albums:  []models.Album{},
			Tracks:  []models.Track{},
		})
		return
	}
	results := h.Store.Search(query)
	writeJSON(w, http.StatusOK, results)
}

// --- Browse ---

// Browse handles /api/browse endpoints
func (h *Handler) Browse(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	path := r.URL.Path

	switch {
	case strings.HasSuffix(path, "/genres"):
		genres := h.Store.GetGenres()
		if genres == nil {
			genres = []string{}
		}
		writeJSON(w, http.StatusOK, genres)

	case strings.HasSuffix(path, "/genre"):
		genre := r.URL.Query().Get("name")
		if genre == "" {
			writeError(w, http.StatusBadRequest, "name query param required")
			return
		}
		tracks := h.Store.GetByGenre(genre)
		if tracks == nil {
			tracks = []models.Track{}
		}
		writeJSON(w, http.StatusOK, tracks)

	case strings.HasSuffix(path, "/new-releases"):
		limitStr := r.URL.Query().Get("limit")
		limit := 10
		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil {
				limit = l
			}
		}
		albums := h.Store.GetNewReleases(limit)
		if albums == nil {
			albums = []models.Album{}
		}
		writeJSON(w, http.StatusOK, albums)

	case strings.HasSuffix(path, "/top-tracks"):
		limitStr := r.URL.Query().Get("limit")
		limit := 20
		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil {
				limit = l
			}
		}
		tracks := h.Store.GetTopTracks(limit)
		if tracks == nil {
			tracks = []models.Track{}
		}
		writeJSON(w, http.StatusOK, tracks)

	default:
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"genres":       h.Store.GetGenres(),
			"new_releases": h.Store.GetNewReleases(5),
			"top_tracks":   h.Store.GetTopTracks(10),
		})
	}
}

// --- Queue/Playback ---

// Playback handles /api/playback endpoints
func (h *Handler) Playback(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case strings.HasSuffix(path, "/play"):
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		var req struct {
			TrackID string   `json:"track_id"`
			Queue   []string `json:"queue"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		h.Store.PlayTrack(req.TrackID, req.Queue)
		h.Store.AddRecentlyPlayed(req.TrackID)
		writeJSON(w, http.StatusOK, h.Store.GetPlaybackState())

	case strings.HasSuffix(path, "/pause"):
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		h.Store.TogglePlayPause()
		writeJSON(w, http.StatusOK, h.Store.GetPlaybackState())

	case strings.HasSuffix(path, "/next"):
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		trackID, ok := h.Store.NextTrack()
		if !ok {
			writeError(w, http.StatusBadRequest, "no next track")
			return
		}
		h.Store.AddRecentlyPlayed(trackID)
		writeJSON(w, http.StatusOK, h.Store.GetPlaybackState())

	case strings.HasSuffix(path, "/previous"):
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		trackID, ok := h.Store.PreviousTrack()
		if !ok {
			writeError(w, http.StatusBadRequest, "no previous track")
			return
		}
		h.Store.AddRecentlyPlayed(trackID)
		writeJSON(w, http.StatusOK, h.Store.GetPlaybackState())

	case strings.HasSuffix(path, "/volume"):
		if r.Method != http.MethodPost {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		var req struct {
			Volume int `json:"volume"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid JSON")
			return
		}
		h.Store.SetVolume(req.Volume)
		writeJSON(w, http.StatusOK, h.Store.GetPlaybackState())

	default:
		// GET /api/playback - return current state
		if r.Method != http.MethodGet {
			writeError(w, http.StatusMethodNotAllowed, "method not allowed")
			return
		}
		writeJSON(w, http.StatusOK, h.Store.GetPlaybackState())
	}
}

// --- Stats ---

// StatsHandler handles /api/stats
func (h *Handler) StatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	writeJSON(w, http.StatusOK, h.Store.GetStats())
}
