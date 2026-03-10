# Music Playlist Manager

## What This Is

A sleek, Spotify-inspired music playlist manager web application built with Go (stdlib net/http) and deployed to Vercel. Features a dark, immersive UI for managing artists, albums, tracks, and playlists with in-memory storage and rich seed data.

## Core Value

Users can browse, organize, and simulate playback of music through playlists with a polished, Spotify-like dark UI experience — complete with album art, now-playing animations, and smooth transitions.

## Requirements

### Validated

(None yet — ship to validate)

### Active

- [ ] Backend API with CRUD for artists, albums, tracks, playlists
- [ ] Library management with liked songs and recently played history
- [ ] Search across artists, albums, tracks with fuzzy matching
- [ ] Browse by genre, new releases, top tracks
- [ ] Playback queue with next/previous
- [ ] Listening stats (most played artists, genres, total time)
- [ ] In-memory storage with seed data (10 artists, 20 albums, 80 tracks, 3 playlists)
- [ ] Spotify-inspired dark theme UI (#121212 bg, #181818 cards, #1DB954 accent)
- [ ] Bottom player bar with album art, track info, progress bar, controls
- [ ] Sidebar navigation with playlist list and library sections
- [ ] Album grid with cover art placeholders and hover play overlay
- [ ] Track list with alternating rows, duration, heart/like icon
- [ ] Playlist view with large header (cover, title, track count, duration)
- [ ] Now playing animation (bouncing bars next to current track)
- [ ] Smooth transitions between views (fade in/out)
- [ ] Volume slider
- [ ] Search page with categorized results
- [ ] Artist page with top tracks, albums, related artists
- [ ] Responsive: collapsible sidebar on mobile, bottom nav

### Out of Scope

- Actual audio playback — simulated playback only
- User authentication — single-user app
- Persistent storage — in-memory only
- Real streaming integration (Spotify/Apple Music APIs)
- Mobile native apps

## Context

- Go stdlib net/http for the backend — no frameworks
- Vercel deployment using `api/index.go` entry point with `@vercel/go`
- All frontend is served as embedded HTML/CSS/JS from the Go binary
- Use `pkg/` for packages (not `internal/` — Vercel rejects it)
- Seed data should provide a realistic browsing experience on first load

## Constraints

- **Stack**: Go stdlib only (net/http), no external Go dependencies
- **Deployment**: Vercel serverless functions via @vercel/go
- **Package layout**: Must use `pkg/` not `internal/` (Vercel limitation)
- **Entry point**: `api/index.go` as the Vercel handler
- **Frontend**: Embedded in Go (no separate build step), vanilla HTML/CSS/JS

## Key Decisions

| Decision | Rationale | Outcome |
|----------|-----------|---------|
| Go stdlib net/http | Simplicity, no deps, Vercel compatible | — Pending |
| In-memory storage | No database needed, simpler deployment | — Pending |
| Embedded frontend | Single deployment unit, no build pipeline | — Pending |
| pkg/ over internal/ | Vercel rejects internal/ packages | — Pending |

---
*Last updated: 2026-03-10 after initialization*
