package store

import (
	"fmt"
	"strings"
	"sync"

	"github.com/shootdaj/ax-test-music/pkg/models"
)

// Store is the in-memory data store with thread-safe access
type Store struct {
	mu sync.RWMutex

	artists   map[string]models.Artist
	albums    map[string]models.Album
	tracks    map[string]models.Track
	playlists map[string]models.Playlist

	library  models.Library
	playback models.PlaybackState

	// Stats tracking
	artistPlayCounts map[string]int
	totalListeningMs int64

	// ID counters
	nextID int
}

// New creates a new empty Store
func New() *Store {
	return &Store{
		artists:          make(map[string]models.Artist),
		albums:           make(map[string]models.Album),
		tracks:           make(map[string]models.Track),
		playlists:        make(map[string]models.Playlist),
		library:          models.Library{LikedTrackIDs: []string{}, RecentlyPlayed: []string{}},
		playback:         models.PlaybackState{Queue: []string{}, Volume: 80},
		artistPlayCounts: make(map[string]int),
		nextID:           1,
	}
}

// GenerateID generates a unique ID with the given prefix
func (s *Store) GenerateID(prefix string) string {
	s.nextID++
	return fmt.Sprintf("%s-%d", prefix, s.nextID-1)
}

// --- Artists ---

// CreateArtist adds a new artist
func (s *Store) CreateArtist(a models.Artist) models.Artist {
	s.mu.Lock()
	defer s.mu.Unlock()
	if a.ID == "" {
		a.ID = s.GenerateID("artist")
	}
	s.artists[a.ID] = a
	return a
}

// GetArtist returns an artist by ID
func (s *Store) GetArtist(id string) (models.Artist, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	a, ok := s.artists[id]
	return a, ok
}

// ListArtists returns all artists
func (s *Store) ListArtists() []models.Artist {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]models.Artist, 0, len(s.artists))
	for _, a := range s.artists {
		result = append(result, a)
	}
	return result
}

// UpdateArtist updates an existing artist
func (s *Store) UpdateArtist(a models.Artist) (models.Artist, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.artists[a.ID]; !ok {
		return a, false
	}
	s.artists[a.ID] = a
	return a, true
}

// DeleteArtist removes an artist by ID
func (s *Store) DeleteArtist(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.artists[id]; !ok {
		return false
	}
	delete(s.artists, id)
	return true
}

// --- Albums ---

// CreateAlbum adds a new album
func (s *Store) CreateAlbum(a models.Album) models.Album {
	s.mu.Lock()
	defer s.mu.Unlock()
	if a.ID == "" {
		a.ID = s.GenerateID("album")
	}
	s.albums[a.ID] = a
	return a
}

// GetAlbum returns an album by ID
func (s *Store) GetAlbum(id string) (models.Album, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	a, ok := s.albums[id]
	return a, ok
}

// ListAlbums returns all albums
func (s *Store) ListAlbums() []models.Album {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]models.Album, 0, len(s.albums))
	for _, a := range s.albums {
		result = append(result, a)
	}
	return result
}

// GetAlbumsByArtist returns albums for a specific artist
func (s *Store) GetAlbumsByArtist(artistID string) []models.Album {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []models.Album
	for _, a := range s.albums {
		if a.ArtistID == artistID {
			result = append(result, a)
		}
	}
	return result
}

// UpdateAlbum updates an existing album
func (s *Store) UpdateAlbum(a models.Album) (models.Album, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.albums[a.ID]; !ok {
		return a, false
	}
	s.albums[a.ID] = a
	return a, true
}

// DeleteAlbum removes an album by ID
func (s *Store) DeleteAlbum(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.albums[id]; !ok {
		return false
	}
	delete(s.albums, id)
	return true
}

// --- Tracks ---

// CreateTrack adds a new track
func (s *Store) CreateTrack(t models.Track) models.Track {
	s.mu.Lock()
	defer s.mu.Unlock()
	if t.ID == "" {
		t.ID = s.GenerateID("track")
	}
	s.tracks[t.ID] = t
	return t
}

// GetTrack returns a track by ID
func (s *Store) GetTrack(id string) (models.Track, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.tracks[id]
	return t, ok
}

// ListTracks returns all tracks
func (s *Store) ListTracks() []models.Track {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]models.Track, 0, len(s.tracks))
	for _, t := range s.tracks {
		result = append(result, t)
	}
	return result
}

// GetTracksByAlbum returns tracks for a specific album
func (s *Store) GetTracksByAlbum(albumID string) []models.Track {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []models.Track
	for _, t := range s.tracks {
		if t.AlbumID == albumID {
			result = append(result, t)
		}
	}
	return result
}

// GetTracksByArtist returns tracks for a specific artist
func (s *Store) GetTracksByArtist(artistID string) []models.Track {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []models.Track
	for _, t := range s.tracks {
		if t.ArtistID == artistID {
			result = append(result, t)
		}
	}
	return result
}

// UpdateTrack updates an existing track
func (s *Store) UpdateTrack(t models.Track) (models.Track, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tracks[t.ID]; !ok {
		return t, false
	}
	s.tracks[t.ID] = t
	return t, true
}

// DeleteTrack removes a track by ID
func (s *Store) DeleteTrack(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tracks[id]; !ok {
		return false
	}
	delete(s.tracks, id)
	return true
}

// --- Playlists ---

// CreatePlaylist adds a new playlist
func (s *Store) CreatePlaylist(p models.Playlist) models.Playlist {
	s.mu.Lock()
	defer s.mu.Unlock()
	if p.ID == "" {
		p.ID = s.GenerateID("playlist")
	}
	if p.TrackIDs == nil {
		p.TrackIDs = []string{}
	}
	s.playlists[p.ID] = p
	return p
}

// GetPlaylist returns a playlist by ID
func (s *Store) GetPlaylist(id string) (models.Playlist, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p, ok := s.playlists[id]
	return p, ok
}

// ListPlaylists returns all playlists
func (s *Store) ListPlaylists() []models.Playlist {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]models.Playlist, 0, len(s.playlists))
	for _, p := range s.playlists {
		result = append(result, p)
	}
	return result
}

// UpdatePlaylist updates an existing playlist
func (s *Store) UpdatePlaylist(p models.Playlist) (models.Playlist, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.playlists[p.ID]; !ok {
		return p, false
	}
	s.playlists[p.ID] = p
	return p, true
}

// DeletePlaylist removes a playlist by ID
func (s *Store) DeletePlaylist(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.playlists[id]; !ok {
		return false
	}
	delete(s.playlists, id)
	return true
}

// AddTrackToPlaylist adds a track to a playlist
func (s *Store) AddTrackToPlaylist(playlistID, trackID string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.playlists[playlistID]
	if !ok {
		return false
	}
	// Check if track exists
	if _, ok := s.tracks[trackID]; !ok {
		return false
	}
	p.TrackIDs = append(p.TrackIDs, trackID)
	s.playlists[playlistID] = p
	return true
}

// RemoveTrackFromPlaylist removes a track from a playlist by index
func (s *Store) RemoveTrackFromPlaylist(playlistID string, index int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.playlists[playlistID]
	if !ok {
		return false
	}
	if index < 0 || index >= len(p.TrackIDs) {
		return false
	}
	p.TrackIDs = append(p.TrackIDs[:index], p.TrackIDs[index+1:]...)
	s.playlists[playlistID] = p
	return true
}

// ReorderPlaylistTrack moves a track from one position to another
func (s *Store) ReorderPlaylistTrack(playlistID string, fromIndex, toIndex int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, ok := s.playlists[playlistID]
	if !ok {
		return false
	}
	if fromIndex < 0 || fromIndex >= len(p.TrackIDs) || toIndex < 0 || toIndex >= len(p.TrackIDs) {
		return false
	}
	trackID := p.TrackIDs[fromIndex]
	// Remove from old position
	p.TrackIDs = append(p.TrackIDs[:fromIndex], p.TrackIDs[fromIndex+1:]...)
	// Insert at new position
	newTracks := make([]string, 0, len(p.TrackIDs)+1)
	newTracks = append(newTracks, p.TrackIDs[:toIndex]...)
	newTracks = append(newTracks, trackID)
	newTracks = append(newTracks, p.TrackIDs[toIndex:]...)
	p.TrackIDs = newTracks
	s.playlists[playlistID] = p
	return true
}

// --- Library ---

// ToggleLike toggles the liked state of a track
func (s *Store) ToggleLike(trackID string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.tracks[trackID]; !ok {
		return false
	}
	for i, id := range s.library.LikedTrackIDs {
		if id == trackID {
			s.library.LikedTrackIDs = append(s.library.LikedTrackIDs[:i], s.library.LikedTrackIDs[i+1:]...)
			return true
		}
	}
	s.library.LikedTrackIDs = append(s.library.LikedTrackIDs, trackID)
	return true
}

// IsLiked checks if a track is liked
func (s *Store) IsLiked(trackID string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, id := range s.library.LikedTrackIDs {
		if id == trackID {
			return true
		}
	}
	return false
}

// GetLikedTracks returns all liked track IDs
func (s *Store) GetLikedTracks() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]string, len(s.library.LikedTrackIDs))
	copy(result, s.library.LikedTrackIDs)
	return result
}

// AddRecentlyPlayed adds a track to recently played history
func (s *Store) AddRecentlyPlayed(trackID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Remove if already exists
	for i, id := range s.library.RecentlyPlayed {
		if id == trackID {
			s.library.RecentlyPlayed = append(s.library.RecentlyPlayed[:i], s.library.RecentlyPlayed[i+1:]...)
			break
		}
	}
	// Add to front
	s.library.RecentlyPlayed = append([]string{trackID}, s.library.RecentlyPlayed...)
	// Keep max 50
	if len(s.library.RecentlyPlayed) > 50 {
		s.library.RecentlyPlayed = s.library.RecentlyPlayed[:50]
	}
}

// GetRecentlyPlayed returns recently played track IDs
func (s *Store) GetRecentlyPlayed() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make([]string, len(s.library.RecentlyPlayed))
	copy(result, s.library.RecentlyPlayed)
	return result
}

// --- Search ---

// Search searches across artists, albums, and tracks by name
func (s *Store) Search(query string) models.SearchResults {
	s.mu.RLock()
	defer s.mu.RUnlock()
	q := strings.ToLower(query)
	results := models.SearchResults{
		Artists: []models.Artist{},
		Albums:  []models.Album{},
		Tracks:  []models.Track{},
	}
	for _, a := range s.artists {
		if fuzzyMatch(a.Name, q) || fuzzyMatch(a.Genre, q) {
			results.Artists = append(results.Artists, a)
		}
	}
	for _, a := range s.albums {
		if fuzzyMatch(a.Title, q) || fuzzyMatch(a.Genre, q) {
			results.Albums = append(results.Albums, a)
		}
	}
	for _, t := range s.tracks {
		if fuzzyMatch(t.Title, q) || fuzzyMatch(t.Genre, q) {
			results.Tracks = append(results.Tracks, t)
		}
	}
	return results
}

// fuzzyMatch checks if the target contains the query (case-insensitive)
func fuzzyMatch(target, query string) bool {
	target = strings.ToLower(target)
	// Direct substring match
	if strings.Contains(target, query) {
		return true
	}
	// Simple fuzzy: check if all query chars appear in order
	qi := 0
	for _, c := range target {
		if qi < len(query) && byte(c) == query[qi] {
			qi++
		}
	}
	return qi == len(query)
}

// --- Browse ---

// GetByGenre returns tracks filtered by genre
func (s *Store) GetByGenre(genre string) []models.Track {
	s.mu.RLock()
	defer s.mu.RUnlock()
	g := strings.ToLower(genre)
	var result []models.Track
	for _, t := range s.tracks {
		if strings.ToLower(t.Genre) == g {
			result = append(result, t)
		}
	}
	return result
}

// GetGenres returns all unique genres
func (s *Store) GetGenres() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	genreMap := make(map[string]bool)
	for _, t := range s.tracks {
		if t.Genre != "" {
			genreMap[t.Genre] = true
		}
	}
	for _, a := range s.artists {
		if a.Genre != "" {
			genreMap[a.Genre] = true
		}
	}
	result := make([]string, 0, len(genreMap))
	for g := range genreMap {
		result = append(result, g)
	}
	return result
}

// GetNewReleases returns the most recently added albums (by highest ID)
func (s *Store) GetNewReleases(limit int) []models.Album {
	s.mu.RLock()
	defer s.mu.RUnlock()
	albums := make([]models.Album, 0, len(s.albums))
	for _, a := range s.albums {
		albums = append(albums, a)
	}
	// Sort by year descending (newer first)
	for i := 0; i < len(albums); i++ {
		for j := i + 1; j < len(albums); j++ {
			if albums[j].Year > albums[i].Year {
				albums[i], albums[j] = albums[j], albums[i]
			}
		}
	}
	if limit > 0 && limit < len(albums) {
		albums = albums[:limit]
	}
	return albums
}

// GetTopTracks returns tracks that appear in the most playlists
func (s *Store) GetTopTracks(limit int) []models.Track {
	s.mu.RLock()
	defer s.mu.RUnlock()
	trackCount := make(map[string]int)
	for _, p := range s.playlists {
		for _, tid := range p.TrackIDs {
			trackCount[tid]++
		}
	}
	type trackScore struct {
		track models.Track
		count int
	}
	var scored []trackScore
	for tid, count := range trackCount {
		if t, ok := s.tracks[tid]; ok {
			scored = append(scored, trackScore{track: t, count: count})
		}
	}
	// Sort by count descending
	for i := 0; i < len(scored); i++ {
		for j := i + 1; j < len(scored); j++ {
			if scored[j].count > scored[i].count {
				scored[i], scored[j] = scored[j], scored[i]
			}
		}
	}
	result := make([]models.Track, 0, limit)
	for i, s := range scored {
		if limit > 0 && i >= limit {
			break
		}
		result = append(result, s.track)
	}
	return result
}

// --- Playback ---

// GetPlaybackState returns current playback state
func (s *Store) GetPlaybackState() models.PlaybackState {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.playback
}

// SetPlaybackState sets the playback state
func (s *Store) SetPlaybackState(state models.PlaybackState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.playback = state
}

// PlayTrack sets a track as currently playing
func (s *Store) PlayTrack(trackID string, queue []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.playback.CurrentTrackID = trackID
	s.playback.IsPlaying = true
	s.playback.ProgressMs = 0
	if queue != nil {
		s.playback.Queue = queue
		// Find position in queue
		for i, id := range queue {
			if id == trackID {
				s.playback.QueuePosition = i
				break
			}
		}
	}
	// Track play count
	if t, ok := s.tracks[trackID]; ok {
		s.artistPlayCounts[t.ArtistID]++
		s.totalListeningMs += int64(t.DurationSeconds) * 1000
	}
}

// NextTrack moves to the next track in the queue
func (s *Store) NextTrack() (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.playback.Queue) == 0 {
		return "", false
	}
	nextPos := s.playback.QueuePosition + 1
	if nextPos >= len(s.playback.Queue) {
		return "", false
	}
	s.playback.QueuePosition = nextPos
	s.playback.CurrentTrackID = s.playback.Queue[nextPos]
	s.playback.ProgressMs = 0
	// Track play count
	if t, ok := s.tracks[s.playback.CurrentTrackID]; ok {
		s.artistPlayCounts[t.ArtistID]++
		s.totalListeningMs += int64(t.DurationSeconds) * 1000
	}
	return s.playback.CurrentTrackID, true
}

// PreviousTrack moves to the previous track in the queue
func (s *Store) PreviousTrack() (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.playback.Queue) == 0 {
		return "", false
	}
	prevPos := s.playback.QueuePosition - 1
	if prevPos < 0 {
		return "", false
	}
	s.playback.QueuePosition = prevPos
	s.playback.CurrentTrackID = s.playback.Queue[prevPos]
	s.playback.ProgressMs = 0
	return s.playback.CurrentTrackID, true
}

// TogglePlayPause toggles the play/pause state
func (s *Store) TogglePlayPause() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.playback.IsPlaying = !s.playback.IsPlaying
	return s.playback.IsPlaying
}

// SetVolume sets the volume level (0-100)
func (s *Store) SetVolume(vol int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if vol < 0 {
		vol = 0
	}
	if vol > 100 {
		vol = 100
	}
	s.playback.Volume = vol
}

// --- Stats ---

// GetStats returns listening statistics
func (s *Store) GetStats() models.Stats {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Build most played artists
	type artistCount struct {
		id    string
		count int
	}
	var counts []artistCount
	for id, count := range s.artistPlayCounts {
		counts = append(counts, artistCount{id: id, count: count})
	}
	// Sort descending
	for i := 0; i < len(counts); i++ {
		for j := i + 1; j < len(counts); j++ {
			if counts[j].count > counts[i].count {
				counts[i], counts[j] = counts[j], counts[i]
			}
		}
	}
	var mostPlayed []models.ArtistPlayCount
	for i, ac := range counts {
		if i >= 10 {
			break
		}
		name := ac.id
		if a, ok := s.artists[ac.id]; ok {
			name = a.Name
		}
		mostPlayed = append(mostPlayed, models.ArtistPlayCount{
			ArtistID:  ac.id,
			Name:      name,
			PlayCount: ac.count,
		})
	}

	// Genre distribution
	genreDist := make(map[string]int)
	for _, t := range s.tracks {
		if t.Genre != "" {
			genreDist[t.Genre]++
		}
	}

	return models.Stats{
		MostPlayedArtists: mostPlayed,
		GenreDistribution: genreDist,
		TotalListeningMs:  s.totalListeningMs,
	}
}
