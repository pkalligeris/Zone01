# ASCII Art Web

Web application that converts input strings into ASCII art using selectable banner templates.

## Description

ASCII Art Web is a web-based version of the ASCII art generator. It provides a graphical user interface where users can input text and select from three different banner styles (standard, shadow, thinkertoy) to generate ASCII art. The application is built with Go's standard library using net/http for the server and html/template for rendering.

## Authors

- panokatos

## Usage

### How to Run

Start the web server:

```bash
go run ./cmd/ascii-art-web
```

The server will start on `http://localhost:8080`

Open your web browser and navigate to:
```
http://localhost:8080
```

### Using the Application

1. Enter your text in the text input field
2. Select a banner style (standard, shadow, or thinkertoy) using the radio buttons
3. Click "Generate ASCII Art" button
4. The ASCII art result will be displayed below the form on the same page

## Implementation Details

### Algorithm

The web application follows this flow:

1. **GET /** - Serves the main HTML page with an empty form
   - Loads the HTML template from `templates/index.html`
   - Displays text input field and banner selection radio buttons
   - Returns HTTP 200 on success, 404 for invalid paths, 500 for template errors

2. **POST /ascii-art** - Processes the form submission
   - Validates that text and banner fields are present (400 if missing)
   - Validates banner selection is one of: standard, shadow, thinkertoy (400 if invalid)
   - Validates input contains only ASCII characters 32-126 plus newlines (400 if invalid)
   - Normalizes line endings (CRLF → LF)
   - Loads the selected banner file from `assets/banners/` (404 if not found)
   - Parses banner into character map (500 if malformed)
   - Renders ASCII art using existing render engine
   - Returns result on the same page with preserved form values
   - Returns appropriate HTTP status codes: 200 (success), 400 (bad request), 404 (not found), 500 (server error)

3. **Rendering Engine** - Reuses existing CLI logic
   - Each ASCII character is 8 lines tall
   - Characters are concatenated horizontally
   - Newlines create separate 8-line blocks
   - Uses the banner map (rune → []string) from `internal/banner`
   - Rendering logic from `internal/render`

### Architecture

```
cmd/ascii-art-web/main.go    → Web server entry point
templates/index.html          → HTML template with form
internal/banner/              → Banner loading (reused from CLI)
internal/render/              → ASCII rendering (reused from CLI)
pkg/model/                    → Shared data structures
assets/banners/               → Banner font files
```

### HTTP Status Codes

- **200 OK** - Successful request (GET / or POST /ascii-art with valid data)
- **400 Bad Request** - Invalid input (missing fields, invalid banner, invalid characters)
- **404 Not Found** - Invalid route or banner file not found
- **405 Method Not Allowed** - Wrong HTTP method used
- **500 Internal Server Error** - Template loading error, banner parsing error, or rendering error

### Template System

Uses Go's `html/template` package with a single template file that:
- Displays the input form
- Preserves user input and banner selection after POST
- Shows error messages in red box if validation fails
- Shows ASCII art result in monospace pre-formatted block
- Auto-escapes HTML to prevent XSS attacks
