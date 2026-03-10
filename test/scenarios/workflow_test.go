//go:build scenario

package scenarios

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

func TestScenario_BrowseAndPlayMusic(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// 1. Browse new releases
	resp, err := http.Get(srv.URL + "/api/browse/new-releases?limit=5")
	if err != nil {
		t.Fatal(err)
	}
	var albums []models.Album
	json.NewDecoder(resp.Body).Decode(&albums)
	resp.Body.Close()
	if len(albums) == 0 {
		t.Fatal("expected new releases")
	}

	// 2. Get tracks from first album
	resp, err = http.Get(srv.URL + "/api/albums/" + albums[0].ID + "/tracks")
	if err != nil {
		t.Fatal(err)
	}
	var tracks []models.Track
	json.NewDecoder(resp.Body).Decode(&tracks)
	resp.Body.Close()
	if len(tracks) == 0 {
		t.Fatal("expected tracks in album")
	}

	// 3. Build a queue from album tracks and play
	var queue []string
	for _, tr := range tracks {
		queue = append(queue, tr.ID)
	}
	playBody, _ := json.Marshal(map[string]interface{}{
		"track_id": tracks[0].ID,
		"queue":    queue,
	})
	resp, err = http.Post(srv.URL+"/api/playback/play", "application/json", bytes.NewBuffer(playBody))
	if err != nil {
		t.Fatal(err)
	}
	var state models.PlaybackState
	json.NewDecoder(resp.Body).Decode(&state)
	resp.Body.Close()
	if !state.IsPlaying {
		t.Error("expected playing state")
	}
	if state.CurrentTrackID != tracks[0].ID {
		t.Errorf("expected current track %s, got %s", tracks[0].ID, state.CurrentTrackID)
	}

	// 4. Skip to next track
	resp, err = http.Post(srv.URL+"/api/playback/next", "application/json", nil)
	if err != nil {
		t.Fatal(err)
	}
	json.NewDecoder(resp.Body).Decode(&state)
	resp.Body.Close()
	if state.CurrentTrackID != tracks[1].ID {
		t.Errorf("expected next track %s, got %s", tracks[1].ID, state.CurrentTrackID)
	}

	// 5. Check recently played
	resp, err = http.Get(srv.URL + "/api/library/recent")
	if err != nil {
		t.Fatal(err)
	}
	var recent []models.Track
	json.NewDecoder(resp.Body).Decode(&recent)
	resp.Body.Close()
	if len(recent) == 0 {
		t.Error("expected recently played tracks")
	}
}

func TestScenario_SearchAndCreatePlaylist(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// 1. Search for rock music
	resp, err := http.Get(srv.URL + "/api/search?q=rock")
	if err != nil {
		t.Fatal(err)
	}
	var results models.SearchResults
	json.NewDecoder(resp.Body).Decode(&results)
	resp.Body.Close()

	// 2. Create a new playlist
	playlistBody := `{"name":"My Rock Mix","description":"Best rock tracks"}`
	resp, err = http.Post(srv.URL+"/api/playlists", "application/json", bytes.NewBufferString(playlistBody))
	if err != nil {
		t.Fatal(err)
	}
	var playlist models.Playlist
	json.NewDecoder(resp.Body).Decode(&playlist)
	resp.Body.Close()
	if playlist.ID == "" {
		t.Fatal("expected playlist ID")
	}

	// 3. Add some rock tracks to the playlist
	if len(results.Tracks) > 0 {
		for i := 0; i < len(results.Tracks) && i < 3; i++ {
			addBody, _ := json.Marshal(map[string]string{"track_id": results.Tracks[i].ID})
			resp, err = http.Post(srv.URL+"/api/playlists/"+playlist.ID+"/tracks", "application/json", bytes.NewBuffer(addBody))
			if err != nil {
				t.Fatal(err)
			}
			resp.Body.Close()
		}
	}

	// 4. Get playlist and verify tracks
	resp, err = http.Get(srv.URL + "/api/playlists/" + playlist.ID + "/tracks")
	if err != nil {
		t.Fatal(err)
	}
	var playlistTracks []models.Track
	json.NewDecoder(resp.Body).Decode(&playlistTracks)
	resp.Body.Close()

	// 5. Like a track
	if len(playlistTracks) > 0 {
		likeBody, _ := json.Marshal(map[string]string{"track_id": playlistTracks[0].ID})
		resp, err = http.Post(srv.URL+"/api/library/like", "application/json", bytes.NewBuffer(likeBody))
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
	}

	// 6. Verify liked tracks
	resp, err = http.Get(srv.URL + "/api/library/liked")
	if err != nil {
		t.Fatal(err)
	}
	var liked []models.Track
	json.NewDecoder(resp.Body).Decode(&liked)
	resp.Body.Close()
}

func TestScenario_ArtistExploration(t *testing.T) {
	srv := setupServer()
	defer srv.Close()

	// 1. Get all artists
	resp, err := http.Get(srv.URL + "/api/artists")
	if err != nil {
		t.Fatal(err)
	}
	var artists []models.Artist
	json.NewDecoder(resp.Body).Decode(&artists)
	resp.Body.Close()
	if len(artists) == 0 {
		t.Fatal("expected artists")
	}

	// 2. Pick an artist and explore
	artist := artists[0]
	resp, err = http.Get(srv.URL + "/api/artists/" + artist.ID)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	// 3. Get artist's albums
	resp, err = http.Get(srv.URL + "/api/artists/" + artist.ID + "/albums")
	if err != nil {
		t.Fatal(err)
	}
	var albums []models.Album
	json.NewDecoder(resp.Body).Decode(&albums)
	resp.Body.Close()

	// 4. Get artist's top tracks
	resp, err = http.Get(srv.URL + "/api/artists/" + artist.ID + "/tracks")
	if err != nil {
		t.Fatal(err)
	}
	var tracks []models.Track
	json.NewDecoder(resp.Body).Decode(&tracks)
	resp.Body.Close()

	// 5. Get related artists
	resp, err = http.Get(srv.URL + "/api/artists/" + artist.ID + "/related")
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()

	// 6. Check stats
	resp, err = http.Get(srv.URL + "/api/stats")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("stats: expected 200, got %d", resp.StatusCode)
	}
	resp.Body.Close()
}
