# Task 14: Ascii-Art Web Export

**Objective:** Extend the web application to support exporting the generated ASCII art as a downloadable text file with the appropriate HTTP headers and permissions.

## Steps

1.  **Define Export Endpoint**
    *   Create a new HTTP GET or POST endpoint (e.g., `/export`) to handle the file download.
    *   The endpoint should receive the necessary text and configuration (banner, etc.) either from query parameters, a form submission, or a re-render from session data, to generate the art. Or, allow exporting directly from the displayed result.

2.  **Generate ASCII Art for Export**
    *   Reuse the shared processing logic (input validation, banner loading, rendering) to generate the raw ASCII string.
    *   Ensure the output respects the `standard`, `shadow`, or `thinkertoy` banner formats.

3.  **Implement HTTP Headers for File Transfer**
    *   Set `Content-Type: text/plain` to specify the file format.
    *   Set `Content-Length` matching the bytes of the generated ASCII art.
    *   Set `Content-Disposition: attachment; filename="ascii_art.txt"` to trigger a file download instead of displaying in the browser.

4.  **Handle Export Errors Gracefully**
    *   If the text is invalid, banner missing, or render fails, return an appropriate HTTP status code (e.g., `400 Bad Request`, `404 Not Found`, `500 Internal Server Error`).
    *   Ensure the user receives a readable error page or response instead of a broken download.

5.  **Update the Web Interface**
    *   Add a "Download" or "Export as TXT" button to `templates/index.html`.
    *   Wire the button to trigger the export endpoint with the current input parameters.

6.  **Code Quality & Best Practices**
    *   Only use standard Go packages (`net/http`, `strings`, `strconv`, etc.).
    *   Ensure the generated file has read/write permissions for the user upon download (browser default behavior handles local permissions, but server response should be clean).

## Acceptance Criteria
*   [ ] The web interface includes a button/link to download the ASCII art.
*   [ ] Clicking the download button triggers a file download named `ascii_art.txt` (or similar).
*   [ ] The downloaded file correctly contains the generated ASCII art in plain text.
*   [ ] The server includes `Content-Type`, `Content-Length`, and `Content-Disposition` headers in the export response.
*   [ ] Invalid requests (bad characters, missing banner) display an error instead of downloading an empty/broken file.
*   [ ] Only standard Go packages are used for the implementation.
