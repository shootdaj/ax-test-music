package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shootdaj/ax-test-music/pkg/models"
	"github.com/shootdaj/ax-test-music/pkg/seed"
	"github.com/shootdaj/ax-test-music/pkg/store"
)

func setupHandler() *Handler {
	s := store.New()
	seed.LoadSeedData(s)
	return New(s)
}

func TestHealth(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	w := httptest.NewRecorder()
	h.Health(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestArtists_List(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/artists", nil)
	w := httptest.NewRecorder()
	h.Artists(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var artists []models.Artist
	json.NewDecoder(w.Body).Decode(&artists)
	if len(artists) != 10 {
		t.Errorf("expected 10 seed artists, got %d", len(artists))
	}
}

func TestArtists_GetByID(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/artists/artist-1", nil)
	w := httptest.NewRecorder()
	h.Artists(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var artist models.Artist
	json.NewDecoder(w.Body).Decode(&artist)
	if artist.Name != "Luna Eclipse" {
		t.Errorf("expected 'Luna Eclipse', got %q", artist.Name)
	}
}

func TestArtists_GetNotFound(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/artists/nonexistent", nil)
	w := httptest.NewRecorder()
	h.Artists(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}

func TestArtists_Create(t *testing.T) {
	h := setupHandler()
	body := `{"name":"New Artist","genre":"Pop"}`
	req := httptest.NewRequest(http.MethodPost, "/api/artists", bytes.NewBufferString(body))
	w := httptest.NewRecorder()
	h.Artists(w, req)
	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
}

func TestArtists_Update(t *testing.T) {
	h := setupHandler()
	body := `{"name":"Updated Luna"}`
	req := httptest.NewRequest(http.MethodPut, "/api/artists/artist-1", bytes.NewBufferString(body))
	w := httptest.NewRecorder()
	h.Artists(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestArtists_Delete(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodDelete, "/api/artists/artist-1", nil)
	w := httptest.NewRecorder()
	h.Artists(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestAlbums_List(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/albums", nil)
	w := httptest.NewRecorder()
	h.Albums(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var albums []models.Album
	json.NewDecoder(w.Body).Decode(&albums)
	if len(albums) != 20 {
		t.Errorf("expected 20 seed albums, got %d", len(albums))
	}
}

func TestAlbums_GetByID(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/albums/album-1", nil)
	w := httptest.NewRecorder()
	h.Albums(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestTracks_List(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/tracks", nil)
	w := httptest.NewRecorder()
	h.Tracks(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var tracks []models.Track
	json.NewDecoder(w.Body).Decode(&tracks)
	if len(tracks) != 80 {
		t.Errorf("expected 80 seed tracks, got %d", len(tracks))
	}
}

func TestPlaylists_List(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/playlists", nil)
	w := httptest.NewRecorder()
	h.Playlists(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var playlists []models.Playlist
	json.NewDecoder(w.Body).Decode(&playlists)
	if len(playlists) != 3 {
		t.Errorf("expected 3 seed playlists, got %d", len(playlists))
	}
}

func TestPlaylists_GetByID(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/playlists/playlist-1", nil)
	w := httptest.NewRecorder()
	h.Playlists(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestPlaylistTracks_AddAndGet(t *testing.T) {
	h := setupHandler()
	// Add track to playlist
	body := `{"track_id":"track-1"}`
	req := httptest.NewRequest(http.MethodPost, "/api/playlists/playlist-1/tracks", bytes.NewBufferString(body))
	w := httptest.NewRecorder()
	h.Playlists(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	// Get playlist tracks
	req = httptest.NewRequest(http.MethodGet, "/api/playlists/playlist-1/tracks", nil)
	w = httptest.NewRecorder()
	h.Playlists(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestSearch(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/search?q=luna", nil)
	w := httptest.NewRecorder()
	h.SearchHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	var results models.SearchResults
	json.NewDecoder(w.Body).Decode(&results)
	if len(results.Artists) == 0 {
		t.Error("expected to find artist 'Luna Eclipse'")
	}
}

func TestSearch_Empty(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/search?q=", nil)
	w := httptest.NewRecorder()
	h.SearchHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestBrowse_Genres(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/browse/genres", nil)
	w := httptest.NewRecorder()
	h.Browse(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestBrowse_NewReleases(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/browse/new-releases", nil)
	w := httptest.NewRecorder()
	h.Browse(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestBrowse_TopTracks(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/browse/top-tracks", nil)
	w := httptest.NewRecorder()
	h.Browse(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestLibrary_Like(t *testing.T) {
	h := setupHandler()
	body := `{"track_id":"track-1"}`
	req := httptest.NewRequest(http.MethodPost, "/api/library/like", bytes.NewBufferString(body))
	w := httptest.NewRecorder()
	h.Library(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestLibrary_GetLiked(t *testing.T) {
	h := setupHandler()
	// Like a track first
	body := `{"track_id":"track-1"}`
	req := httptest.NewRequest(http.MethodPost, "/api/library/like", bytes.NewBufferString(body))
	w := httptest.NewRecorder()
	h.Library(w, req)

	// Get liked
	req = httptest.NewRequest(http.MethodGet, "/api/library/liked", nil)
	w = httptest.NewRecorder()
	h.Library(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestPlayback_Play(t *testing.T) {
	h := setupHandler()
	body := `{"track_id":"track-1","queue":["track-1","track-2","track-3"]}`
	req := httptest.NewRequest(http.MethodPost, "/api/playback/play", bytes.NewBufferString(body))
	w := httptest.NewRecorder()
	h.Playback(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestPlayback_GetState(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/playback", nil)
	w := httptest.NewRecorder()
	h.Playback(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestStats(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/api/stats", nil)
	w := httptest.NewRecorder()
	h.StatsHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestFrontend(t *testing.T) {
	h := setupHandler()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	h.ServeFrontend(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); ct != "text/html; charset=utf-8" {
		t.Errorf("expected text/html content type, got %q", ct)
	}
}
