//go:build integration

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shootdaj/ax-test-music/pkg/handlers"
	"github.com/shootdaj/ax-test-music/pkg/models"
	"github.com/shootdaj/ax-test-music/pkg/router"
	"github.com/shootdaj/ax-test-music/pkg/seed"
	"github.com/shootdaj/ax-test-music/pkg/store"
)

func setupServer() *httptest.Server {
	s := store.New()
	seed.LoadSeedData(s)
	h := handlers.New(s)
	r := router.New(h)
	return httptest.NewServer(r)
}

func TestAPI_HealthEndpoint(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/api/health")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestAPI_ArtistsCRUD(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// List
	resp, err := http.Get(srv.URL + "/api/artists")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("list: expected 200, got %d", resp.StatusCode)
	}
	var artists []models.Artist
	json.NewDecoder(resp.Body).Decode(&artists)
	resp.Body.Close()
	if len(artists) != 10 {
		t.Errorf("expected 10 artists, got %d", len(artists))
	}

	// Create
	body := `{"name":"Integration Artist","genre":"Test"}`
	resp, err = http.Post(srv.URL+"/api/artists", "application/json", bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 201 {
		t.Errorf("create: expected 201, got %d", resp.StatusCode)
	}
	var created models.Artist
	json.NewDecoder(resp.Body).Decode(&created)
	resp.Body.Close()

	// Get
	resp, err = http.Get(srv.URL + "/api/artists/" + created.ID)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("get: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// Delete
	req, _ := http.NewRequest(http.MethodDelete, srv.URL+"/api/artists/"+created.ID, nil)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("delete: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()
}

func TestAPI_AlbumsCRUD(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// List
	resp, err := http.Get(srv.URL + "/api/albums")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	var albums []models.Album
	json.NewDecoder(resp.Body).Decode(&albums)
	resp.Body.Close()
	if len(albums) != 20 {
		t.Errorf("expected 20 albums, got %d", len(albums))
	}
}

func TestAPI_TracksCRUD(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/api/tracks")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	var tracks []models.Track
	json.NewDecoder(resp.Body).Decode(&tracks)
	resp.Body.Close()
	if len(tracks) != 80 {
		t.Errorf("expected 80 tracks, got %d", len(tracks))
	}
}

func TestAPI_PlaylistsCRUD(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// List
	resp, err := http.Get(srv.URL + "/api/playlists")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	var playlists []models.Playlist
	json.NewDecoder(resp.Body).Decode(&playlists)
	resp.Body.Close()
	if len(playlists) != 3 {
		t.Errorf("expected 3 playlists, got %d", len(playlists))
	}
}

func TestAPI_PlaylistTrackManagement(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// Add track to playlist
	body := `{"track_id":"track-5"}`
	resp, err := http.Post(srv.URL+"/api/playlists/playlist-1/tracks", "application/json", bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("add track: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// Get playlist tracks
	resp, err = http.Get(srv.URL + "/api/playlists/playlist-1/tracks")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("get tracks: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()
}

func TestAPI_Search(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/api/search?q=luna")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	var results models.SearchResults
	json.NewDecoder(resp.Body).Decode(&results)
	resp.Body.Close()
	if len(results.Artists) == 0 {
		t.Error("expected to find artist matching 'luna'")
	}
}

func TestAPI_Browse(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// Genres
	resp, err := http.Get(srv.URL + "/api/browse/genres")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("genres: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// New releases
	resp, err = http.Get(srv.URL + "/api/browse/new-releases")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("new-releases: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// Top tracks
	resp, err = http.Get(srv.URL + "/api/browse/top-tracks")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("top-tracks: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()
}

func TestAPI_Library(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// Like a track
	body := `{"track_id":"track-1"}`
	resp, err := http.Post(srv.URL+"/api/library/like", "application/json", bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("like: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// Get liked
	resp, err = http.Get(srv.URL + "/api/library/liked")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("liked: expected 200, got %d", resp.StatusCode)
	}
	var liked []models.Track
	json.NewDecoder(resp.Body).Decode(&liked)
	resp.Body.Close()
	if len(liked) != 1 {
		t.Errorf("expected 1 liked track, got %d", len(liked))
	}
}

func TestAPI_Playback(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// Play a track
	body := `{"track_id":"track-1","queue":["track-1","track-2","track-3"]}`
	resp, err := http.Post(srv.URL+"/api/playback/play", "application/json", bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("play: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()

	// Get state
	resp, err = http.Get(srv.URL + "/api/playback")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("state: expected 200, got %d", resp.StatusCode)
	}
	var state models.PlaybackState
	json.NewDecoder(resp.Body).Decode(&state)
	resp.Body.Close()
	if state.CurrentTrackID != "track-1" {
		t.Errorf("expected current track track-1, got %s", state.CurrentTrackID)
	}

	// Next
	resp, err = http.Post(srv.URL+"/api/playback/next", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("next: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()
}

func TestAPI_Stats(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/api/stats")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()
}

func TestAPI_Frontend(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html; charset=utf-8" {
		t.Errorf("expected text/html, got %q", ct)
	}
	resp.Body.Close()
}
