package models

// Artist represents a music artist
type Artist struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	ImageURL string `json:"image_url"`
	Bio      string `json:"bio"`
}

// Album represents a music album
type Album struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ArtistID string `json:"artist_id"`
	Year     int    `json:"year"`
	CoverURL string `json:"cover_url"`
	Genre    string `json:"genre"`
}

// Track represents a single track
type Track struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	AlbumID         string `json:"album_id"`
	ArtistID        string `json:"artist_id"`
	DurationSeconds int    `json:"duration_seconds"`
	TrackNumber     int    `json:"track_number"`
	Genre           string `json:"genre"`
}

// Playlist represents a user playlist
type Playlist struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	CoverURL    string   `json:"cover_url"`
	TrackIDs    []string `json:"track_ids"`
}

// PlaybackState represents the current playback state
type PlaybackState struct {
	CurrentTrackID string   `json:"current_track_id"`
	Queue          []string `json:"queue"`
	QueuePosition  int      `json:"queue_position"`
	IsPlaying      bool     `json:"is_playing"`
	ProgressMs     int64    `json:"progress_ms"`
	Volume         int      `json:"volume"`
}

// Library represents the user's library
type Library struct {
	LikedTrackIDs  []string `json:"liked_track_ids"`
	RecentlyPlayed []string `json:"recently_played"`
}

// Stats represents listening statistics
type Stats struct {
	MostPlayedArtists []ArtistPlayCount `json:"most_played_artists"`
	GenreDistribution map[string]int    `json:"genre_distribution"`
	TotalListeningMs  int64             `json:"total_listening_ms"`
}

// ArtistPlayCount tracks play count per artist
type ArtistPlayCount struct {
	ArtistID  string `json:"artist_id"`
	Name      string `json:"name"`
	PlayCount int    `json:"play_count"`
}

// SearchResults holds categorized search results
type SearchResults struct {
	Artists []Artist `json:"artists"`
	Albums  []Album  `json:"albums"`
	Tracks  []Track  `json:"tracks"`
}
