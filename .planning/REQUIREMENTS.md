# Requirements: Music Playlist Manager

**Defined:** 2026-03-10
**Core Value:** Users can browse, organize, and simulate playback of music through playlists with a polished, Spotify-like dark UI experience

## v1 Requirements

### Data Models & API

- [ ] **API-01**: CRUD endpoints for artists (name, genre, image_url, bio)
- [ ] **API-02**: CRUD endpoints for albums (title, artist_id, year, cover_url, genre)
- [ ] **API-03**: CRUD endpoints for tracks (title, album_id, artist_id, duration_seconds, track_number, genre)
- [ ] **API-04**: CRUD endpoints for playlists (name, description, cover_url)
- [ ] **API-05**: Add/remove/reorder tracks in playlists
- [ ] **API-06**: In-memory storage with thread-safe access

### Library

- [ ] **LIB-01**: User can like/unlike songs (liked songs collection)
- [ ] **LIB-02**: Recently played history tracking
- [ ] **LIB-03**: Library view showing liked songs and recent history

### Search & Browse

- [ ] **SRCH-01**: Search across artists, albums, tracks by name with fuzzy matching
- [ ] **SRCH-02**: Browse by genre with filterable results
- [ ] **SRCH-03**: New releases section (recently added albums/tracks)
- [ ] **SRCH-04**: Top tracks section (most added to playlists)

### Queue & Playback

- [ ] **PLAY-01**: Current playback queue management
- [ ] **PLAY-02**: Next/previous track navigation
- [ ] **PLAY-03**: Play/pause state management
- [ ] **PLAY-04**: Simulated progress bar (time-based, no real audio)

### Stats

- [ ] **STAT-01**: Most played artists statistics
- [ ] **STAT-02**: Genre distribution statistics
- [ ] **STAT-03**: Total listening time calculation

### Seed Data

- [ ] **SEED-01**: 10 artists with varied genres
- [ ] **SEED-02**: 20 albums distributed across artists
- [ ] **SEED-03**: 80 tracks distributed across albums
- [ ] **SEED-04**: 3 pre-built playlists with tracks

### Frontend - Layout

- [ ] **UI-01**: Spotify-inspired dark theme (#121212 bg, #181818 cards, #1DB954 accent)
- [ ] **UI-02**: Bottom player bar with album art thumbnail, track name/artist, progress bar, play/pause/skip
- [ ] **UI-03**: Sidebar navigation with playlist list and library sections
- [ ] **UI-04**: Responsive design: collapsible sidebar on mobile, bottom nav replaces sidebar

### Frontend - Views

- [ ] **VIEW-01**: Home/browse page with genre sections, new releases, top tracks
- [ ] **VIEW-02**: Album grid view with cover art placeholders (colored squares with album initials), hover play button overlay
- [ ] **VIEW-03**: Track list with alternating row shading, duration, heart/like icon
- [ ] **VIEW-04**: Playlist view with large header (cover + title + track count + total duration)
- [ ] **VIEW-05**: Search page with categorized results (artists, albums, tracks in separate sections)
- [ ] **VIEW-06**: Artist page with top tracks, albums, related artists
- [ ] **VIEW-07**: Library page showing liked songs and recently played

### Frontend - Interactions

- [ ] **INTR-01**: Now playing animation (bouncing bars icon next to currently playing track)
- [ ] **INTR-02**: Smooth transitions between views (fade in/out)
- [ ] **INTR-03**: Volume slider with sleek design
- [ ] **INTR-04**: Hover effects on album cards and track rows

### Deployment

- [ ] **DEPL-01**: Vercel deployment via api/index.go entry point
- [ ] **DEPL-02**: vercel.json with correct routing configuration
- [ ] **DEPL-03**: go.mod at project root

## v2 Requirements

### Enhanced Features

- **ENH-01**: Drag-and-drop playlist reordering
- **ENH-02**: Playlist cover image upload
- **ENH-03**: Social sharing of playlists
- **ENH-04**: Keyboard shortcuts for playback control

## Out of Scope

| Feature | Reason |
|---------|--------|
| Real audio playback | Simulated playback only — no audio files |
| User authentication | Single-user app, no auth needed |
| Persistent database | In-memory storage for simplicity |
| Spotify/Apple Music API integration | Standalone app with seed data |
| Mobile native apps | Web-only, responsive design |

## Traceability

| Requirement | Phase | Status |
|-------------|-------|--------|
| API-01 | Phase 1 | Pending |
| API-02 | Phase 1 | Pending |
| API-03 | Phase 1 | Pending |
| API-04 | Phase 1 | Pending |
| API-05 | Phase 1 | Pending |
| API-06 | Phase 1 | Pending |
| SEED-01 | Phase 1 | Pending |
| SEED-02 | Phase 1 | Pending |
| SEED-03 | Phase 1 | Pending |
| SEED-04 | Phase 1 | Pending |
| DEPL-01 | Phase 1 | Pending |
| DEPL-02 | Phase 1 | Pending |
| DEPL-03 | Phase 1 | Pending |
| LIB-01 | Phase 2 | Pending |
| LIB-02 | Phase 2 | Pending |
| LIB-03 | Phase 2 | Pending |
| SRCH-01 | Phase 2 | Pending |
| SRCH-02 | Phase 2 | Pending |
| SRCH-03 | Phase 2 | Pending |
| SRCH-04 | Phase 2 | Pending |
| PLAY-01 | Phase 3 | Pending |
| PLAY-02 | Phase 3 | Pending |
| PLAY-03 | Phase 3 | Pending |
| PLAY-04 | Phase 3 | Pending |
| STAT-01 | Phase 3 | Pending |
| STAT-02 | Phase 3 | Pending |
| STAT-03 | Phase 3 | Pending |
| UI-01 | Phase 4 | Pending |
| UI-02 | Phase 4 | Pending |
| UI-03 | Phase 4 | Pending |
| UI-04 | Phase 4 | Pending |
| VIEW-01 | Phase 4 | Pending |
| VIEW-02 | Phase 4 | Pending |
| VIEW-03 | Phase 4 | Pending |
| VIEW-04 | Phase 4 | Pending |
| VIEW-05 | Phase 4 | Pending |
| VIEW-06 | Phase 4 | Pending |
| VIEW-07 | Phase 4 | Pending |
| INTR-01 | Phase 5 | Pending |
| INTR-02 | Phase 5 | Pending |
| INTR-03 | Phase 5 | Pending |
| INTR-04 | Phase 5 | Pending |

**Coverage:**
- v1 requirements: 40 total
- Mapped to phases: 40
- Unmapped: 0

---
*Requirements defined: 2026-03-10*
*Last updated: 2026-03-10 after initial definition*
