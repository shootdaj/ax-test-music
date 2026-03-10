# Music Playlist Manager

## Project Overview
Spotify-inspired music playlist manager built with Go stdlib net/http, deployed to Vercel.

## Stack
- Go stdlib net/http (no frameworks)
- Vercel deployment via api/index.go
- In-memory storage
- Vanilla HTML/CSS/JS frontend embedded in Go

## Key Conventions
- Use `pkg/` for packages (NOT `internal/` - Vercel rejects it)
- Entry point: `api/index.go`
- Dark theme: #121212 bg, #181818 cards, #1DB954 accent

# Testing Requirements (AX)

Every feature implementation MUST include tests at all three tiers:

## Test Tiers
1. **Unit tests** -- Test individual functions/methods in isolation. Mock external dependencies.
2. **Integration tests** -- Test component interactions with real services via docker-compose.test.yml.
3. **Scenario tests** -- Test full user workflows end-to-end.

## Test Naming
Use semantic names: `Test<Component>_<Behavior>[_<Condition>]`
- Good: `TestAuthService_LoginWithValidCredentials`, `TestFullCheckoutFlow`
- Bad: `TestShouldWork`, `Test1`, `TestGivenUserWhenLoginThenSuccess`

## Reference
- See `TEST_GUIDE.md` for requirement-to-test mapping
- See `.claude/ax/references/testing-pyramid.md` for full methodology
- Every requirement in ROADMAP.md must map to at least one scenario test
