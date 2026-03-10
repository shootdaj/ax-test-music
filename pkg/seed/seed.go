package seed

import (
	"github.com/shootdaj/ax-test-music/pkg/models"
	"github.com/shootdaj/ax-test-music/pkg/store"
)

// LoadSeedData populates the store with sample music data
func LoadSeedData(s *store.Store) {
	// 10 Artists
	artists := []models.Artist{
		{ID: "artist-1", Name: "Luna Eclipse", Genre: "Electronic", ImageURL: "", Bio: "Pioneering electronic artist known for ethereal soundscapes and pulsating beats."},
		{ID: "artist-2", Name: "The Midnight Riders", Genre: "Rock", ImageURL: "", Bio: "Classic rock band with a modern edge, known for powerful guitar riffs."},
		{ID: "artist-3", Name: "Aria Santos", Genre: "Pop", ImageURL: "", Bio: "Chart-topping pop vocalist with soulful melodies and catchy hooks."},
		{ID: "artist-4", Name: "Jazz Collective", Genre: "Jazz", ImageURL: "", Bio: "An ensemble of world-class jazz musicians pushing the boundaries of improvisation."},
		{ID: "artist-5", Name: "Neon Dreams", Genre: "Synthwave", ImageURL: "", Bio: "Retro-futuristic synthwave duo creating nostalgic 80s-inspired electronic music."},
		{ID: "artist-6", Name: "Rhythm Nation", Genre: "Hip Hop", ImageURL: "", Bio: "Groundbreaking hip hop group blending conscious lyrics with innovative production."},
		{ID: "artist-7", Name: "Acoustic Hearts", Genre: "Folk", ImageURL: "", Bio: "Indie folk band crafting intimate stories through acoustic arrangements."},
		{ID: "artist-8", Name: "Bassline Theory", Genre: "Electronic", ImageURL: "", Bio: "Deep house and techno producer known for hypnotic basslines."},
		{ID: "artist-9", Name: "Sierra Wave", Genre: "R&B", ImageURL: "", Bio: "Smooth R&B artist with a voice that blends contemporary and classic soul."},
		{ID: "artist-10", Name: "Thunder Valley", Genre: "Metal", ImageURL: "", Bio: "Progressive metal band known for complex compositions and thunderous performances."},
	}
	for _, a := range artists {
		s.CreateArtist(a)
	}

	// 20 Albums
	albums := []models.Album{
		{ID: "album-1", Title: "Stellar Waves", ArtistID: "artist-1", Year: 2024, CoverURL: "", Genre: "Electronic"},
		{ID: "album-2", Title: "Digital Dawn", ArtistID: "artist-1", Year: 2023, CoverURL: "", Genre: "Electronic"},
		{ID: "album-3", Title: "Highway Anthems", ArtistID: "artist-2", Year: 2024, CoverURL: "", Genre: "Rock"},
		{ID: "album-4", Title: "Rebel Hearts", ArtistID: "artist-2", Year: 2022, CoverURL: "", Genre: "Rock"},
		{ID: "album-5", Title: "Golden Hour", ArtistID: "artist-3", Year: 2024, CoverURL: "", Genre: "Pop"},
		{ID: "album-6", Title: "Midnight Bloom", ArtistID: "artist-3", Year: 2023, CoverURL: "", Genre: "Pop"},
		{ID: "album-7", Title: "Blue Notes", ArtistID: "artist-4", Year: 2024, CoverURL: "", Genre: "Jazz"},
		{ID: "album-8", Title: "Smoky Sessions", ArtistID: "artist-4", Year: 2022, CoverURL: "", Genre: "Jazz"},
		{ID: "album-9", Title: "Retrograde", ArtistID: "artist-5", Year: 2024, CoverURL: "", Genre: "Synthwave"},
		{ID: "album-10", Title: "Chrome Horizons", ArtistID: "artist-5", Year: 2023, CoverURL: "", Genre: "Synthwave"},
		{ID: "album-11", Title: "Street Poetry", ArtistID: "artist-6", Year: 2024, CoverURL: "", Genre: "Hip Hop"},
		{ID: "album-12", Title: "Urban Canvas", ArtistID: "artist-6", Year: 2023, CoverURL: "", Genre: "Hip Hop"},
		{ID: "album-13", Title: "Campfire Stories", ArtistID: "artist-7", Year: 2024, CoverURL: "", Genre: "Folk"},
		{ID: "album-14", Title: "Wanderlust", ArtistID: "artist-7", Year: 2022, CoverURL: "", Genre: "Folk"},
		{ID: "album-15", Title: "Deep Currents", ArtistID: "artist-8", Year: 2024, CoverURL: "", Genre: "Electronic"},
		{ID: "album-16", Title: "Subterranean", ArtistID: "artist-8", Year: 2023, CoverURL: "", Genre: "Electronic"},
		{ID: "album-17", Title: "Velvet Nights", ArtistID: "artist-9", Year: 2024, CoverURL: "", Genre: "R&B"},
		{ID: "album-18", Title: "Silk & Soul", ArtistID: "artist-9", Year: 2023, CoverURL: "", Genre: "R&B"},
		{ID: "album-19", Title: "Iron Forge", ArtistID: "artist-10", Year: 2024, CoverURL: "", Genre: "Metal"},
		{ID: "album-20", Title: "Storm Chaser", ArtistID: "artist-10", Year: 2022, CoverURL: "", Genre: "Metal"},
	}
	for _, a := range albums {
		s.CreateAlbum(a)
	}

	// 80 Tracks (4 per album)
	tracks := []models.Track{
		// Album 1: Stellar Waves (Electronic - Luna Eclipse)
		{ID: "track-1", Title: "Cosmic Drift", AlbumID: "album-1", ArtistID: "artist-1", DurationSeconds: 245, TrackNumber: 1, Genre: "Electronic"},
		{ID: "track-2", Title: "Nebula Rising", AlbumID: "album-1", ArtistID: "artist-1", DurationSeconds: 312, TrackNumber: 2, Genre: "Electronic"},
		{ID: "track-3", Title: "Starlight Pulse", AlbumID: "album-1", ArtistID: "artist-1", DurationSeconds: 278, TrackNumber: 3, Genre: "Electronic"},
		{ID: "track-4", Title: "Event Horizon", AlbumID: "album-1", ArtistID: "artist-1", DurationSeconds: 356, TrackNumber: 4, Genre: "Electronic"},
		// Album 2: Digital Dawn (Electronic - Luna Eclipse)
		{ID: "track-5", Title: "Binary Sunrise", AlbumID: "album-2", ArtistID: "artist-1", DurationSeconds: 223, TrackNumber: 1, Genre: "Electronic"},
		{ID: "track-6", Title: "Pixel Rain", AlbumID: "album-2", ArtistID: "artist-1", DurationSeconds: 267, TrackNumber: 2, Genre: "Electronic"},
		{ID: "track-7", Title: "Circuit Breaker", AlbumID: "album-2", ArtistID: "artist-1", DurationSeconds: 298, TrackNumber: 3, Genre: "Electronic"},
		{ID: "track-8", Title: "Digital Echoes", AlbumID: "album-2", ArtistID: "artist-1", DurationSeconds: 334, TrackNumber: 4, Genre: "Electronic"},
		// Album 3: Highway Anthems (Rock - The Midnight Riders)
		{ID: "track-9", Title: "Open Road", AlbumID: "album-3", ArtistID: "artist-2", DurationSeconds: 234, TrackNumber: 1, Genre: "Rock"},
		{ID: "track-10", Title: "Thunder Rolling", AlbumID: "album-3", ArtistID: "artist-2", DurationSeconds: 276, TrackNumber: 2, Genre: "Rock"},
		{ID: "track-11", Title: "Burning Bridges", AlbumID: "album-3", ArtistID: "artist-2", DurationSeconds: 301, TrackNumber: 3, Genre: "Rock"},
		{ID: "track-12", Title: "Midnight Drive", AlbumID: "album-3", ArtistID: "artist-2", DurationSeconds: 258, TrackNumber: 4, Genre: "Rock"},
		// Album 4: Rebel Hearts (Rock - The Midnight Riders)
		{ID: "track-13", Title: "Wildfire", AlbumID: "album-4", ArtistID: "artist-2", DurationSeconds: 245, TrackNumber: 1, Genre: "Rock"},
		{ID: "track-14", Title: "Steel & Stone", AlbumID: "album-4", ArtistID: "artist-2", DurationSeconds: 289, TrackNumber: 2, Genre: "Rock"},
		{ID: "track-15", Title: "Outlaw Blues", AlbumID: "album-4", ArtistID: "artist-2", DurationSeconds: 312, TrackNumber: 3, Genre: "Rock"},
		{ID: "track-16", Title: "Last Stand", AlbumID: "album-4", ArtistID: "artist-2", DurationSeconds: 267, TrackNumber: 4, Genre: "Rock"},
		// Album 5: Golden Hour (Pop - Aria Santos)
		{ID: "track-17", Title: "Sunlit Dreams", AlbumID: "album-5", ArtistID: "artist-3", DurationSeconds: 198, TrackNumber: 1, Genre: "Pop"},
		{ID: "track-18", Title: "Dancing in Color", AlbumID: "album-5", ArtistID: "artist-3", DurationSeconds: 212, TrackNumber: 2, Genre: "Pop"},
		{ID: "track-19", Title: "Gravity Pull", AlbumID: "album-5", ArtistID: "artist-3", DurationSeconds: 234, TrackNumber: 3, Genre: "Pop"},
		{ID: "track-20", Title: "Golden Light", AlbumID: "album-5", ArtistID: "artist-3", DurationSeconds: 245, TrackNumber: 4, Genre: "Pop"},
		// Album 6: Midnight Bloom (Pop - Aria Santos)
		{ID: "track-21", Title: "Moonflower", AlbumID: "album-6", ArtistID: "artist-3", DurationSeconds: 223, TrackNumber: 1, Genre: "Pop"},
		{ID: "track-22", Title: "Petal Storm", AlbumID: "album-6", ArtistID: "artist-3", DurationSeconds: 198, TrackNumber: 2, Genre: "Pop"},
		{ID: "track-23", Title: "Night Garden", AlbumID: "album-6", ArtistID: "artist-3", DurationSeconds: 256, TrackNumber: 3, Genre: "Pop"},
		{ID: "track-24", Title: "Bloom", AlbumID: "album-6", ArtistID: "artist-3", DurationSeconds: 278, TrackNumber: 4, Genre: "Pop"},
		// Album 7: Blue Notes (Jazz - Jazz Collective)
		{ID: "track-25", Title: "Sapphire Mood", AlbumID: "album-7", ArtistID: "artist-4", DurationSeconds: 345, TrackNumber: 1, Genre: "Jazz"},
		{ID: "track-26", Title: "Velvet Swing", AlbumID: "album-7", ArtistID: "artist-4", DurationSeconds: 423, TrackNumber: 2, Genre: "Jazz"},
		{ID: "track-27", Title: "Cool Breeze", AlbumID: "album-7", ArtistID: "artist-4", DurationSeconds: 312, TrackNumber: 3, Genre: "Jazz"},
		{ID: "track-28", Title: "Indigo Sky", AlbumID: "album-7", ArtistID: "artist-4", DurationSeconds: 378, TrackNumber: 4, Genre: "Jazz"},
		// Album 8: Smoky Sessions (Jazz - Jazz Collective)
		{ID: "track-29", Title: "Late Night Blues", AlbumID: "album-8", ArtistID: "artist-4", DurationSeconds: 389, TrackNumber: 1, Genre: "Jazz"},
		{ID: "track-30", Title: "Whiskey Sour", AlbumID: "album-8", ArtistID: "artist-4", DurationSeconds: 356, TrackNumber: 2, Genre: "Jazz"},
		{ID: "track-31", Title: "Dim Lights", AlbumID: "album-8", ArtistID: "artist-4", DurationSeconds: 298, TrackNumber: 3, Genre: "Jazz"},
		{ID: "track-32", Title: "Closing Time", AlbumID: "album-8", ArtistID: "artist-4", DurationSeconds: 412, TrackNumber: 4, Genre: "Jazz"},
		// Album 9: Retrograde (Synthwave - Neon Dreams)
		{ID: "track-33", Title: "Neon City", AlbumID: "album-9", ArtistID: "artist-5", DurationSeconds: 267, TrackNumber: 1, Genre: "Synthwave"},
		{ID: "track-34", Title: "Laser Grid", AlbumID: "album-9", ArtistID: "artist-5", DurationSeconds: 234, TrackNumber: 2, Genre: "Synthwave"},
		{ID: "track-35", Title: "Chrome Dreams", AlbumID: "album-9", ArtistID: "artist-5", DurationSeconds: 289, TrackNumber: 3, Genre: "Synthwave"},
		{ID: "track-36", Title: "Retrowave", AlbumID: "album-9", ArtistID: "artist-5", DurationSeconds: 312, TrackNumber: 4, Genre: "Synthwave"},
		// Album 10: Chrome Horizons (Synthwave - Neon Dreams)
		{ID: "track-37", Title: "Sunset Drive", AlbumID: "album-10", ArtistID: "artist-5", DurationSeconds: 245, TrackNumber: 1, Genre: "Synthwave"},
		{ID: "track-38", Title: "Electric Skyline", AlbumID: "album-10", ArtistID: "artist-5", DurationSeconds: 278, TrackNumber: 2, Genre: "Synthwave"},
		{ID: "track-39", Title: "VHS Memories", AlbumID: "album-10", ArtistID: "artist-5", DurationSeconds: 223, TrackNumber: 3, Genre: "Synthwave"},
		{ID: "track-40", Title: "Horizon Pulse", AlbumID: "album-10", ArtistID: "artist-5", DurationSeconds: 298, TrackNumber: 4, Genre: "Synthwave"},
		// Album 11: Street Poetry (Hip Hop - Rhythm Nation)
		{ID: "track-41", Title: "City Lights", AlbumID: "album-11", ArtistID: "artist-6", DurationSeconds: 212, TrackNumber: 1, Genre: "Hip Hop"},
		{ID: "track-42", Title: "Block Party", AlbumID: "album-11", ArtistID: "artist-6", DurationSeconds: 198, TrackNumber: 2, Genre: "Hip Hop"},
		{ID: "track-43", Title: "Real Talk", AlbumID: "album-11", ArtistID: "artist-6", DurationSeconds: 234, TrackNumber: 3, Genre: "Hip Hop"},
		{ID: "track-44", Title: "Crown Heights", AlbumID: "album-11", ArtistID: "artist-6", DurationSeconds: 256, TrackNumber: 4, Genre: "Hip Hop"},
		// Album 12: Urban Canvas (Hip Hop - Rhythm Nation)
		{ID: "track-45", Title: "Paint the Town", AlbumID: "album-12", ArtistID: "artist-6", DurationSeconds: 223, TrackNumber: 1, Genre: "Hip Hop"},
		{ID: "track-46", Title: "Graffiti Soul", AlbumID: "album-12", ArtistID: "artist-6", DurationSeconds: 245, TrackNumber: 2, Genre: "Hip Hop"},
		{ID: "track-47", Title: "Concrete Jungle", AlbumID: "album-12", ArtistID: "artist-6", DurationSeconds: 267, TrackNumber: 3, Genre: "Hip Hop"},
		{ID: "track-48", Title: "Skyline Flow", AlbumID: "album-12", ArtistID: "artist-6", DurationSeconds: 234, TrackNumber: 4, Genre: "Hip Hop"},
		// Album 13: Campfire Stories (Folk - Acoustic Hearts)
		{ID: "track-49", Title: "Morning Trail", AlbumID: "album-13", ArtistID: "artist-7", DurationSeconds: 267, TrackNumber: 1, Genre: "Folk"},
		{ID: "track-50", Title: "River Song", AlbumID: "album-13", ArtistID: "artist-7", DurationSeconds: 234, TrackNumber: 2, Genre: "Folk"},
		{ID: "track-51", Title: "Ember Glow", AlbumID: "album-13", ArtistID: "artist-7", DurationSeconds: 289, TrackNumber: 3, Genre: "Folk"},
		{ID: "track-52", Title: "Starlit Path", AlbumID: "album-13", ArtistID: "artist-7", DurationSeconds: 312, TrackNumber: 4, Genre: "Folk"},
		// Album 14: Wanderlust (Folk - Acoustic Hearts)
		{ID: "track-53", Title: "Open Fields", AlbumID: "album-14", ArtistID: "artist-7", DurationSeconds: 245, TrackNumber: 1, Genre: "Folk"},
		{ID: "track-54", Title: "Mountain Echo", AlbumID: "album-14", ArtistID: "artist-7", DurationSeconds: 278, TrackNumber: 2, Genre: "Folk"},
		{ID: "track-55", Title: "Dusty Road", AlbumID: "album-14", ArtistID: "artist-7", DurationSeconds: 223, TrackNumber: 3, Genre: "Folk"},
		{ID: "track-56", Title: "Homeward Bound", AlbumID: "album-14", ArtistID: "artist-7", DurationSeconds: 298, TrackNumber: 4, Genre: "Folk"},
		// Album 15: Deep Currents (Electronic - Bassline Theory)
		{ID: "track-57", Title: "Undertow", AlbumID: "album-15", ArtistID: "artist-8", DurationSeconds: 334, TrackNumber: 1, Genre: "Electronic"},
		{ID: "track-58", Title: "Pressure Wave", AlbumID: "album-15", ArtistID: "artist-8", DurationSeconds: 356, TrackNumber: 2, Genre: "Electronic"},
		{ID: "track-59", Title: "Deep Dive", AlbumID: "album-15", ArtistID: "artist-8", DurationSeconds: 312, TrackNumber: 3, Genre: "Electronic"},
		{ID: "track-60", Title: "Abyssal Zone", AlbumID: "album-15", ArtistID: "artist-8", DurationSeconds: 378, TrackNumber: 4, Genre: "Electronic"},
		// Album 16: Subterranean (Electronic - Bassline Theory)
		{ID: "track-61", Title: "Cave System", AlbumID: "album-16", ArtistID: "artist-8", DurationSeconds: 289, TrackNumber: 1, Genre: "Electronic"},
		{ID: "track-62", Title: "Echo Chamber", AlbumID: "album-16", ArtistID: "artist-8", DurationSeconds: 312, TrackNumber: 2, Genre: "Electronic"},
		{ID: "track-63", Title: "Underground", AlbumID: "album-16", ArtistID: "artist-8", DurationSeconds: 345, TrackNumber: 3, Genre: "Electronic"},
		{ID: "track-64", Title: "Tunnel Vision", AlbumID: "album-16", ArtistID: "artist-8", DurationSeconds: 267, TrackNumber: 4, Genre: "Electronic"},
		// Album 17: Velvet Nights (R&B - Sierra Wave)
		{ID: "track-65", Title: "Satin Dreams", AlbumID: "album-17", ArtistID: "artist-9", DurationSeconds: 234, TrackNumber: 1, Genre: "R&B"},
		{ID: "track-66", Title: "Candlelight", AlbumID: "album-17", ArtistID: "artist-9", DurationSeconds: 256, TrackNumber: 2, Genre: "R&B"},
		{ID: "track-67", Title: "Slow Dance", AlbumID: "album-17", ArtistID: "artist-9", DurationSeconds: 278, TrackNumber: 3, Genre: "R&B"},
		{ID: "track-68", Title: "After Hours", AlbumID: "album-17", ArtistID: "artist-9", DurationSeconds: 289, TrackNumber: 4, Genre: "R&B"},
		// Album 18: Silk & Soul (R&B - Sierra Wave)
		{ID: "track-69", Title: "Smooth Operator", AlbumID: "album-18", ArtistID: "artist-9", DurationSeconds: 223, TrackNumber: 1, Genre: "R&B"},
		{ID: "track-70", Title: "Soul Fire", AlbumID: "album-18", ArtistID: "artist-9", DurationSeconds: 245, TrackNumber: 2, Genre: "R&B"},
		{ID: "track-71", Title: "Golden Touch", AlbumID: "album-18", ArtistID: "artist-9", DurationSeconds: 267, TrackNumber: 3, Genre: "R&B"},
		{ID: "track-72", Title: "Midnight Serenade", AlbumID: "album-18", ArtistID: "artist-9", DurationSeconds: 312, TrackNumber: 4, Genre: "R&B"},
		// Album 19: Iron Forge (Metal - Thunder Valley)
		{ID: "track-73", Title: "Anvil Strike", AlbumID: "album-19", ArtistID: "artist-10", DurationSeconds: 298, TrackNumber: 1, Genre: "Metal"},
		{ID: "track-74", Title: "Molten Core", AlbumID: "album-19", ArtistID: "artist-10", DurationSeconds: 334, TrackNumber: 2, Genre: "Metal"},
		{ID: "track-75", Title: "Steel Tempest", AlbumID: "album-19", ArtistID: "artist-10", DurationSeconds: 356, TrackNumber: 3, Genre: "Metal"},
		{ID: "track-76", Title: "Forged in Fire", AlbumID: "album-19", ArtistID: "artist-10", DurationSeconds: 378, TrackNumber: 4, Genre: "Metal"},
		// Album 20: Storm Chaser (Metal - Thunder Valley)
		{ID: "track-77", Title: "Lightning Rod", AlbumID: "album-20", ArtistID: "artist-10", DurationSeconds: 267, TrackNumber: 1, Genre: "Metal"},
		{ID: "track-78", Title: "Eye of Storm", AlbumID: "album-20", ArtistID: "artist-10", DurationSeconds: 289, TrackNumber: 2, Genre: "Metal"},
		{ID: "track-79", Title: "Tornado Alley", AlbumID: "album-20", ArtistID: "artist-10", DurationSeconds: 312, TrackNumber: 3, Genre: "Metal"},
		{ID: "track-80", Title: "Hurricane Force", AlbumID: "album-20", ArtistID: "artist-10", DurationSeconds: 345, TrackNumber: 4, Genre: "Metal"},
	}
	for _, t := range tracks {
		s.CreateTrack(t)
	}

	// 3 Playlists
	playlists := []models.Playlist{
		{
			ID:          "playlist-1",
			Name:        "Chill Vibes",
			Description: "Perfect for relaxing evenings",
			CoverURL:    "",
			TrackIDs: []string{
				"track-1", "track-5", "track-25", "track-33", "track-37",
				"track-49", "track-65", "track-69", "track-21", "track-57",
			},
		},
		{
			ID:          "playlist-2",
			Name:        "Energy Boost",
			Description: "High energy tracks to power your day",
			CoverURL:    "",
			TrackIDs: []string{
				"track-9", "track-13", "track-41", "track-73", "track-77",
				"track-34", "track-42", "track-10", "track-74", "track-11",
			},
		},
		{
			ID:          "playlist-3",
			Name:        "Late Night Sessions",
			Description: "Smooth tunes for the midnight hour",
			CoverURL:    "",
			TrackIDs: []string{
				"track-29", "track-30", "track-66", "track-67", "track-61",
				"track-62", "track-23", "track-24", "track-70", "track-71",
			},
		},
	}
	for _, p := range playlists {
		s.CreatePlaylist(p)
	}
}
