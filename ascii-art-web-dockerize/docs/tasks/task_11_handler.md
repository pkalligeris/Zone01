# Task 11: ASCII Art Handler & Input Validation

**Objective:** Implement the `POST /ascii-art` endpoint with full input validation, banner loading, rendering, and error handling.

## Steps

1.  **Implement `asciiArtHandler`**
    *   Accept only `POST`; return `405 Method Not Allowed` otherwise.
    *   Extract `text` and `banner` from form values.

2.  **Input Validation**
    *   Empty text or banner → `400 Bad Request` with error message.
    *   Invalid banner name (not `standard`, `shadow`, or `thinkertoy`) → `400 Bad Request`.
    *   Non-ASCII characters (outside 32–126, excluding `\n` and `\r`) → `400 Bad Request`.

3.  **Processing**
    *   Normalize line endings (`\r\n` → `\n`, `\r` → `\n`).
    *   Load banner via `banner.GetBannerPath` + `banner.LoadBanner`.
    *   Render via `render.Render` with `Config` (left alignment).

4.  **Error Handling**
    *   Banner file not found → `404 Not Found`.
    *   Failed banner load or render → `500 Internal Server Error`.
    *   All errors are displayed in the template alongside the retained user input.

## Acceptance Criteria
*   [x] `POST /ascii-art` generates and displays ASCII art.
*   [x] Non-POST to `/ascii-art` returns `405 Method Not Allowed`.
*   [x] Empty input returns `400 Bad Request` with error message.
*   [x] Invalid banner returns `400 Bad Request` with error message.
*   [x] Non-ASCII characters return `400 Bad Request` with error message.
*   [x] Missing banner file returns `404 Not Found`.
*   [x] Form retains input and banner selection after submission.
