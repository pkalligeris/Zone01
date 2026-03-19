# Task 13: JSON API Endpoint

**Objective:** Extend the existing web server with a JSON API for ASCII art generation, while reusing the current banner loading and rendering pipeline.

## Steps

1.  **Define API Contract**
    *   Add request and response structs for JSON payloads.
    *   Proposed request fields: `text`, `banner`, `align`, `color`, `color_substring`.
    *   Proposed response fields: `result`, `error`.
    *   Standardize content type handling for `application/json`.

2.  **Add API Route**
    *   Register a new route such as `POST /api/ascii-art` in `cmd/ascii-art-web/main.go`.
    *   Accept only `POST`; return `405 Method Not Allowed` for other methods.
    *   Decode and validate JSON request bodies.

3.  **Extract Shared Processing Logic**
    *   Move the common flow into a reusable helper or service:
        *   validate input
        *   normalize line endings
        *   resolve banner path
        *   load banner
        *   build `model.Config`
        *   call `render.Render`
    *   Reuse this logic from both the HTML form handler and the new API handler to avoid duplicate behavior.

4.  **Implement API Validation**
    *   Reject empty `text` or `banner` with `400 Bad Request`.
    *   Reject invalid banner names outside `standard`, `shadow`, `thinkertoy`.
    *   Reject non-ASCII characters outside `32..126`, excluding newline support.
    *   Reject malformed JSON with `400 Bad Request`.
    *   Reject invalid `color` or `align` values with `400 Bad Request`.

5.  **Implement Structured API Errors**
    *   Return JSON error bodies instead of HTML template responses.
    *   Map failures consistently:
        *   invalid input → `400`
        *   wrong method → `405`
        *   missing banner file → `404`
        *   banner parsing or render failure → `500`

6.  **Testing**
    *   Add handler tests for:
        *   successful JSON render
        *   invalid method
        *   malformed JSON
        *   empty input
        *   invalid banner
        *   invalid ASCII input
        *   invalid color/alignment
    *   Verify the API returns JSON with expected status codes and fields.

## Acceptance Criteria
*   [ ] `POST /api/ascii-art` accepts JSON and returns rendered ASCII art in JSON.
*   [ ] The existing HTML interface continues to work without regression.
*   [ ] Shared rendering logic is not duplicated between HTML and API handlers.
*   [ ] API returns `Content-Type: application/json`.
*   [ ] Invalid JSON returns `400 Bad Request` with structured error response.
*   [ ] Invalid banner, color, align, or input text return `400 Bad Request`.
*   [ ] Missing banner file returns `404 Not Found`.
*   [ ] Render or banner parsing failures return `500 Internal Server Error`.
*   [ ] Automated handler tests cover the new API behavior.
