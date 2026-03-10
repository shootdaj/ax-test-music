package handlers

import "net/http"

// ServeFrontend serves the frontend SPA HTML
func (h *Handler) ServeFrontend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(frontendHTML))
}

// frontendHTML is a placeholder - will be replaced with full Spotify-inspired UI in Phase 4
var frontendHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Music Playlist Manager</title>
    <style>
        body {
            margin: 0;
            background: #121212;
            color: #fff;
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
        }
        .container {
            text-align: center;
            padding: 2rem;
        }
        h1 {
            color: #1DB954;
            font-size: 2rem;
            margin-bottom: 0.5rem;
        }
        p {
            color: #b3b3b3;
            font-size: 1.1rem;
        }
        a {
            color: #1DB954;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Music Playlist Manager</h1>
        <p>API is running. Try <a href="/api/health">/api/health</a></p>
        <p>Full UI coming in Phase 4.</p>
    </div>
</body>
</html>`
