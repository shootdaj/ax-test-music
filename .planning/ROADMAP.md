# Roadmap: Music Playlist Manager

**Created:** 2026-03-10
**Phases:** 5
**Requirements covered:** 40/40

## Overview

| # | Phase | Goal | Requirements | Success Criteria |
|---|-------|------|--------------|------------------|
| 1 | Core API & Data Layer | Working REST API with all CRUD endpoints, seed data, and Vercel deployment config | API-01..06, SEED-01..04, DEPL-01..03 | 4 |
| 2 | Library, Search & Browse | Library management, search with fuzzy matching, and browsable content | LIB-01..03, SRCH-01..04 | 4 |
| 3 | Queue, Playback & Stats | Simulated playback queue, playback controls, and listening statistics | PLAY-01..04, STAT-01..03 | 4 |
| 4 | Frontend - Layout & Views | Complete Spotify-inspired UI with all views, dark theme, and responsive layout | UI-01..04, VIEW-01..07 | 5 |
| 5 | Frontend - Interactions & Polish | Now-playing animations, smooth transitions, volume slider, hover effects | INTR-01..04 | 3 |

---

## Phase 1: Core API & Data Layer

**Goal:** Build the complete REST API with CRUD endpoints for artists, albums, tracks, and playlists. Implement in-memory storage with thread-safe access. Populate with rich seed data. Set up Vercel deployment configuration.

**Requirements:** API-01, API-02, API-03, API-04, API-05, API-06, SEED-01, SEED-02, SEED-03, SEED-04, DEPL-01, DEPL-02, DEPL-03

**Success Criteria:**
1. All CRUD endpoints return correct JSON responses (GET/POST/PUT/DELETE for artists, albums, tracks, playlists)
2. Playlist track management works (add, remove, reorder)
3. Seed data loads on startup with 10 artists, 20 albums, 80 tracks, 3 playlists
4. `go test ./...` passes with all API endpoints tested

---

## Phase 2: Library, Search & Browse

**Goal:** Add library management (liked songs, recently played), search with fuzzy matching across all entities, and browse functionality (by genre, new releases, top tracks).

**Requirements:** LIB-01, LIB-02, LIB-03, SRCH-01, SRCH-02, SRCH-03, SRCH-04

**Success Criteria:**
1. Like/unlike songs persists in liked songs collection
2. Search returns relevant results with partial/fuzzy name matching
3. Browse by genre returns filtered results
4. New releases and top tracks endpoints return correct data

---

## Phase 3: Queue, Playback & Stats

**Goal:** Implement playback queue management with next/previous navigation, play/pause state, simulated progress tracking, and listening statistics.

**Requirements:** PLAY-01, PLAY-02, PLAY-03, PLAY-04, STAT-01, STAT-02, STAT-03

**Success Criteria:**
1. Queue can be populated, cleared, and navigated (next/previous)
2. Play/pause state toggles correctly and progress advances
3. Stats endpoints return most played artists, genre distribution, total time
4. All queue and stats operations are thread-safe

---

## Phase 4: Frontend - Layout & Views

**Goal:** Build the complete Spotify-inspired dark UI with sidebar navigation, bottom player bar, and all content views (home, album grid, track list, playlist, search, artist, library). Responsive with collapsible sidebar on mobile.

**Requirements:** UI-01, UI-02, UI-03, UI-04, VIEW-01, VIEW-02, VIEW-03, VIEW-04, VIEW-05, VIEW-06, VIEW-07

**Success Criteria:**
1. Dark theme renders correctly (#121212 bg, #181818 cards, #1DB954 accent)
2. All 7 views load and display data from the API
3. Bottom player bar shows current track info with controls
4. Sidebar navigates between all views
5. Layout adapts to mobile viewport with collapsible sidebar

---

## Phase 5: Frontend - Interactions & Polish

**Goal:** Add polished interactions: now-playing animation, smooth view transitions, volume slider, and hover effects. Final UI polish pass.

**Requirements:** INTR-01, INTR-02, INTR-03, INTR-04

**Success Criteria:**
1. Bouncing bars animation appears next to currently playing track
2. View transitions are smooth (fade in/out, no jarring reloads)
3. Volume slider controls playback volume with sleek visual design

---

## Dependency Graph

```
Phase 1 (API) → Phase 2 (Library/Search) → Phase 3 (Queue/Stats)
                                                    ↓
Phase 1 (API) → Phase 4 (Frontend Views) → Phase 5 (Polish)
```

Phases 2 and 3 build on Phase 1's API. Phase 4 requires API endpoints from Phases 1-3. Phase 5 polishes Phase 4's frontend.
