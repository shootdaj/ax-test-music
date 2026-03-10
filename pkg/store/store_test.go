package store

import (
	"testing"

	"github.com/shootdaj/ax-test-music/pkg/models"
)

func TestStore_CreateAndGetArtist(t *testing.T) {
	s := New()
	artist := s.CreateArtist(models.Artist{Name: "Test Artist", Genre: "Rock"})
	if artist.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	got, ok := s.GetArtist(artist.ID)
	if !ok {
		t.Fatal("expected to find artist")
	}
	if got.Name != "Test Artist" {
		t.Errorf("expected name 'Test Artist', got %q", got.Name)
	}
}

func TestStore_ListArtists(t *testing.T) {
	s := New()
	s.CreateArtist(models.Artist{Name: "A1"})
	s.CreateArtist(models.Artist{Name: "A2"})
	artists := s.ListArtists()
	if len(artists) != 2 {
		t.Errorf("expected 2 artists, got %d", len(artists))
	}
}

func TestStore_UpdateArtist(t *testing.T) {
	s := New()
	artist := s.CreateArtist(models.Artist{Name: "Original"})
	artist.Name = "Updated"
	updated, ok := s.UpdateArtist(artist)
	if !ok {
		t.Fatal("expected update to succeed")
	}
	if updated.Name != "Updated" {
		t.Errorf("expected name 'Updated', got %q", updated.Name)
	}
}

func TestStore_DeleteArtist(t *testing.T) {
	s := New()
	artist := s.CreateArtist(models.Artist{Name: "ToDelete"})
	if !s.DeleteArtist(artist.ID) {
		t.Fatal("expected delete to succeed")
	}
	_, ok := s.GetArtist(artist.ID)
	if ok {
		t.Fatal("expected artist to be deleted")
	}
}

func TestStore_DeleteArtist_NotFound(t *testing.T) {
	s := New()
	if s.DeleteArtist("nonexistent") {
		t.Fatal("expected delete to return false for nonexistent artist")
	}
}

func TestStore_CreateAndGetAlbum(t *testing.T) {
	s := New()
	album := s.CreateAlbum(models.Album{Title: "Test Album", ArtistID: "a1", Year: 2024})
	if album.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	got, ok := s.GetAlbum(album.ID)
	if !ok {
		t.Fatal("expected to find album")
	}
	if got.Title != "Test Album" {
		t.Errorf("expected title 'Test Album', got %q", got.Title)
	}
}

func TestStore_GetAlbumsByArtist(t *testing.T) {
	s := New()
	s.CreateAlbum(models.Album{Title: "A1", ArtistID: "art-1"})
	s.CreateAlbum(models.Album{Title: "A2", ArtistID: "art-1"})
	s.CreateAlbum(models.Album{Title: "A3", ArtistID: "art-2"})
	albums := s.GetAlbumsByArtist("art-1")
	if len(albums) != 2 {
		t.Errorf("expected 2 albums for art-1, got %d", len(albums))
	}
}

func TestStore_CreateAndGetTrack(t *testing.T) {
	s := New()
	track := s.CreateTrack(models.Track{Title: "Test Track", AlbumID: "al1", ArtistID: "a1", DurationSeconds: 200})
	if track.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	got, ok := s.GetTrack(track.ID)
	if !ok {
		t.Fatal("expected to find track")
	}
	if got.Title != "Test Track" {
		t.Errorf("expected title 'Test Track', got %q", got.Title)
	}
}

func TestStore_GetTracksByAlbum(t *testing.T) {
	s := New()
	s.CreateTrack(models.Track{Title: "T1", AlbumID: "al-1", ArtistID: "a1"})
	s.CreateTrack(models.Track{Title: "T2", AlbumID: "al-1", ArtistID: "a1"})
	s.CreateTrack(models.Track{Title: "T3", AlbumID: "al-2", ArtistID: "a1"})
	tracks := s.GetTracksByAlbum("al-1")
	if len(tracks) != 2 {
		t.Errorf("expected 2 tracks for al-1, got %d", len(tracks))
	}
}

func TestStore_CreateAndGetPlaylist(t *testing.T) {
	s := New()
	playlist := s.CreatePlaylist(models.Playlist{Name: "My Playlist"})
	if playlist.ID == "" {
		t.Fatal("expected non-empty ID")
	}
	got, ok := s.GetPlaylist(playlist.ID)
	if !ok {
		t.Fatal("expected to find playlist")
	}
	if got.Name != "My Playlist" {
		t.Errorf("expected name 'My Playlist', got %q", got.Name)
	}
}

func TestStore_AddTrackToPlaylist(t *testing.T) {
	s := New()
	track := s.CreateTrack(models.Track{Title: "T1"})
	playlist := s.CreatePlaylist(models.Playlist{Name: "P1"})
	if !s.AddTrackToPlaylist(playlist.ID, track.ID) {
		t.Fatal("expected add to succeed")
	}
	p, _ := s.GetPlaylist(playlist.ID)
	if len(p.TrackIDs) != 1 || p.TrackIDs[0] != track.ID {
		t.Errorf("expected track in playlist, got %v", p.TrackIDs)
	}
}

func TestStore_RemoveTrackFromPlaylist(t *testing.T) {
	s := New()
	track := s.CreateTrack(models.Track{Title: "T1"})
	playlist := s.CreatePlaylist(models.Playlist{Name: "P1"})
	s.AddTrackToPlaylist(playlist.ID, track.ID)
	if !s.RemoveTrackFromPlaylist(playlist.ID, 0) {
		t.Fatal("expected remove to succeed")
	}
	p, _ := s.GetPlaylist(playlist.ID)
	if len(p.TrackIDs) != 0 {
		t.Errorf("expected empty playlist, got %v", p.TrackIDs)
	}
}

func TestStore_ReorderPlaylistTrack(t *testing.T) {
	s := New()
	t1 := s.CreateTrack(models.Track{Title: "T1"})
	t2 := s.CreateTrack(models.Track{Title: "T2"})
	t3 := s.CreateTrack(models.Track{Title: "T3"})
	playlist := s.CreatePlaylist(models.Playlist{Name: "P1"})
	s.AddTrackToPlaylist(playlist.ID, t1.ID)
	s.AddTrackToPlaylist(playlist.ID, t2.ID)
	s.AddTrackToPlaylist(playlist.ID, t3.ID)
	// Move track 0 to position 2
	if !s.ReorderPlaylistTrack(playlist.ID, 0, 2) {
		t.Fatal("expected reorder to succeed")
	}
	p, _ := s.GetPlaylist(playlist.ID)
	if p.TrackIDs[0] != t2.ID {
		t.Errorf("expected t2 first, got %s", p.TrackIDs[0])
	}
}

func TestStore_ToggleLike(t *testing.T) {
	s := New()
	track := s.CreateTrack(models.Track{Title: "Likeable"})
	// Like
	if !s.ToggleLike(track.ID) {
		t.Fatal("expected toggle to succeed")
	}
	if !s.IsLiked(track.ID) {
		t.Fatal("expected track to be liked")
	}
	// Unlike
	s.ToggleLike(track.ID)
	if s.IsLiked(track.ID) {
		t.Fatal("expected track to be unliked")
	}
}

func TestStore_RecentlyPlayed(t *testing.T) {
	s := New()
	track := s.CreateTrack(models.Track{Title: "Recent"})
	s.AddRecentlyPlayed(track.ID)
	recent := s.GetRecentlyPlayed()
	if len(recent) != 1 || recent[0] != track.ID {
		t.Errorf("expected recently played to contain track, got %v", recent)
	}
}

func TestStore_Search(t *testing.T) {
	s := New()
	s.CreateArtist(models.Artist{ID: "a1", Name: "Luna Eclipse", Genre: "Electronic"})
	s.CreateAlbum(models.Album{ID: "al1", Title: "Stellar Waves", ArtistID: "a1"})
	s.CreateTrack(models.Track{ID: "t1", Title: "Cosmic Drift", AlbumID: "al1", ArtistID: "a1"})

	results := s.Search("luna")
	if len(results.Artists) == 0 {
		t.Error("expected to find artist 'Luna Eclipse'")
	}
	results = s.Search("stellar")
	if len(results.Albums) == 0 {
		t.Error("expected to find album 'Stellar Waves'")
	}
	results = s.Search("cosmic")
	if len(results.Tracks) == 0 {
		t.Error("expected to find track 'Cosmic Drift'")
	}
}

func TestStore_FuzzySearch(t *testing.T) {
	s := New()
	s.CreateArtist(models.Artist{ID: "a1", Name: "Luna Eclipse", Genre: "Electronic"})
	results := s.Search("lna")
	if len(results.Artists) == 0 {
		t.Error("expected fuzzy match for 'lna' -> 'Luna Eclipse'")
	}
}

func TestStore_GetByGenre(t *testing.T) {
	s := New()
	s.CreateTrack(models.Track{ID: "t1", Title: "T1", Genre: "Rock"})
	s.CreateTrack(models.Track{ID: "t2", Title: "T2", Genre: "Rock"})
	s.CreateTrack(models.Track{ID: "t3", Title: "T3", Genre: "Jazz"})
	tracks := s.GetByGenre("Rock")
	if len(tracks) != 2 {
		t.Errorf("expected 2 rock tracks, got %d", len(tracks))
	}
}

func TestStore_GetGenres(t *testing.T) {
	s := New()
	s.CreateTrack(models.Track{ID: "t1", Genre: "Rock"})
	s.CreateTrack(models.Track{ID: "t2", Genre: "Jazz"})
	s.CreateArtist(models.Artist{ID: "a1", Genre: "Pop"})
	genres := s.GetGenres()
	if len(genres) < 3 {
		t.Errorf("expected at least 3 genres, got %d", len(genres))
	}
}

func TestStore_GetTopTracks(t *testing.T) {
	s := New()
	s.CreateTrack(models.Track{ID: "t1", Title: "Popular"})
	s.CreateTrack(models.Track{ID: "t2", Title: "Less Popular"})
	s.CreatePlaylist(models.Playlist{ID: "p1", TrackIDs: []string{"t1", "t2"}})
	s.CreatePlaylist(models.Playlist{ID: "p2", TrackIDs: []string{"t1"}})
	top := s.GetTopTracks(10)
	if len(top) == 0 {
		t.Fatal("expected top tracks")
	}
	if top[0].ID != "t1" {
		t.Errorf("expected most popular track first, got %s", top[0].ID)
	}
}

func TestStore_PlaybackQueue(t *testing.T) {
	s := New()
	s.CreateArtist(models.Artist{ID: "a1", Name: "Test"})
	s.CreateTrack(models.Track{ID: "t1", Title: "Track 1", ArtistID: "a1", DurationSeconds: 200})
	s.CreateTrack(models.Track{ID: "t2", Title: "Track 2", ArtistID: "a1", DurationSeconds: 200})
	s.CreateTrack(models.Track{ID: "t3", Title: "Track 3", ArtistID: "a1", DurationSeconds: 200})

	s.PlayTrack("t1", []string{"t1", "t2", "t3"})
	state := s.GetPlaybackState()
	if state.CurrentTrackID != "t1" {
		t.Errorf("expected current track t1, got %s", state.CurrentTrackID)
	}
	if !state.IsPlaying {
		t.Error("expected playing")
	}

	// Next
	next, ok := s.NextTrack()
	if !ok || next != "t2" {
		t.Errorf("expected next track t2, got %s", next)
	}

	// Previous
	prev, ok := s.PreviousTrack()
	if !ok || prev != "t1" {
		t.Errorf("expected prev track t1, got %s", prev)
	}
}

func TestStore_TogglePlayPause(t *testing.T) {
	s := New()
	s.CreateTrack(models.Track{ID: "t1", ArtistID: "a1", DurationSeconds: 200})
	s.PlayTrack("t1", nil)
	// Should be playing
	if !s.GetPlaybackState().IsPlaying {
		t.Error("expected playing after PlayTrack")
	}
	// Toggle to pause
	playing := s.TogglePlayPause()
	if playing {
		t.Error("expected paused after toggle")
	}
	// Toggle back to play
	playing = s.TogglePlayPause()
	if !playing {
		t.Error("expected playing after second toggle")
	}
}

func TestStore_SetVolume(t *testing.T) {
	s := New()
	s.SetVolume(50)
	if s.GetPlaybackState().Volume != 50 {
		t.Errorf("expected volume 50, got %d", s.GetPlaybackState().Volume)
	}
	s.SetVolume(-10)
	if s.GetPlaybackState().Volume != 0 {
		t.Error("expected volume clamped to 0")
	}
	s.SetVolume(150)
	if s.GetPlaybackState().Volume != 100 {
		t.Error("expected volume clamped to 100")
	}
}

func TestStore_Stats(t *testing.T) {
	s := New()
	s.CreateArtist(models.Artist{ID: "a1", Name: "Artist 1"})
	s.CreateTrack(models.Track{ID: "t1", ArtistID: "a1", DurationSeconds: 200, Genre: "Rock"})
	s.CreateTrack(models.Track{ID: "t2", ArtistID: "a1", DurationSeconds: 300, Genre: "Rock"})
	s.PlayTrack("t1", nil)
	s.PlayTrack("t2", nil)

	stats := s.GetStats()
	if len(stats.MostPlayedArtists) == 0 {
		t.Error("expected most played artists")
	}
	if stats.TotalListeningMs == 0 {
		t.Error("expected non-zero total listening time")
	}
	if stats.GenreDistribution["Rock"] != 2 {
		t.Errorf("expected 2 Rock tracks in distribution, got %d", stats.GenreDistribution["Rock"])
	}
}

func TestStore_NewReleases(t *testing.T) {
	s := New()
	s.CreateAlbum(models.Album{ID: "al1", Title: "Old", Year: 2020})
	s.CreateAlbum(models.Album{ID: "al2", Title: "New", Year: 2024})
	s.CreateAlbum(models.Album{ID: "al3", Title: "Newer", Year: 2025})
	releases := s.GetNewReleases(2)
	if len(releases) != 2 {
		t.Errorf("expected 2 releases, got %d", len(releases))
	}
	if releases[0].Year < releases[1].Year {
		t.Error("expected sorted by year descending")
	}
}

func TestStore_ConcurrentAccess(t *testing.T) {
	s := New()
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			s.CreateArtist(models.Artist{Name: "Concurrent"})
			s.ListArtists()
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	artists := s.ListArtists()
	if len(artists) != 10 {
		t.Errorf("expected 10 artists from concurrent access, got %d", len(artists))
	}
}
