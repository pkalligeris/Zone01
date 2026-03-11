# Task 10: Web Server Setup & Routing

**Objective:** Create the HTTP server entrypoint with route registration and a home page handler, establishing the web infrastructure.

## Steps

1.  **Create Web Entrypoint (`cmd/ascii-art-web/main.go`)**
    *   Define `PageData` struct holding `Input`, `Banner`, `Result`, and `Error` fields.
    *   Parse `templates/index.html` at init time using `html/template`.
    *   Register routes: `GET /`, `POST /ascii-art`, and `/assets/` static file server.
    *   Start HTTP server on `:8080`.

2.  **Implement Home Handler (`homeHandler`)**
    *   Serve `GET /` only; return `404` for any other path.
    *   Render the template with default `Banner: "standard"`.

3.  **Static Asset Serving**
    *   Serve files under `assets/` at the `/assets/` URL path using `http.FileServer`.

## Acceptance Criteria
*   [x] Server starts on port `8080`.
*   [x] `GET /` renders the form with `standard` selected by default.
*   [x] Unknown paths return `404 Not Found`.
*   [x] Static assets are served at `/assets/`.
