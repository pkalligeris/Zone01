# ASCII Art Generator (Go)

ASCII Art Generator is a Go project that converts input strings into banner-based ASCII art. It includes:
- a CLI application
- a browser-based web interface
- a JSON API endpoint for programmatic access

## Features

- Supports ASCII characters `32..126`
- Converts literal `\n` in input into real line breaks
- Preserves case, spacing, and consecutive empty lines
- Supports banner selection: `standard`, `shadow`, `thinkertoy`
- Supports output colorization with `--color=<format>` in the CLI
- Supports output redirection to file via `--output=<file>`
- Supports alignment with `--align=left|center|right|justify`
- Provides a web UI for ASCII art generation
- Provides a JSON API for ASCII art generation
- Includes unit tests and golden integration tests

## Authors

- pkallige
- ipapigki
- nbougial

## Project Layout

- `cmd/ascii-art/main.go`: CLI entrypoint and app wiring
- `cmd/ascii-art-web/main.go`: web server entrypoint
- `internal/web/handlers.go`: HTML and API handlers
- `internal/web/service.go`: shared web validation and render flow
- `internal/web/template.go`: template loading helpers
- `internal/web/types.go`: web request/response and page data types
- `internal/input`: CLI input validation and `\n` parsing
- `internal/banner`: banner file loading/parsing
- `internal/output`: file output writer
- `internal/render`: ASCII rendering engine
- `pkg/model`: shared `Banner` and `Config` types
- `templates/index.html`: web UI template
- `test/integration_test.go`: golden regression suite
- `test/golden/*.txt`: expected outputs
- `Dockerfile`: multi-stage container build for the web server

## Requirements

- Go `1.21+`
- Docker (optional, for containerized web deployment)

## Quick Start

Run tests:

```bash
go test ./...
```

Run CLI:

```bash
go run ./cmd/ascii-art "hello"
```

Run web server:

```bash
go run ./cmd/ascii-art-web
```

Open:

```text
http://localhost:8080
```

## CLI Usage

```bash
go run ./cmd/ascii-art [OPTION] [STRING] [BANNER]
```

**Note:** Flags must use the `=` separator (for example `--color=red`). Space-separated flags are not supported.

| Flag | Description | Format |
|------|-------------|--------|
| `--color` | Colorize the output | `--color=<color>` |
| `--output` | Save output to a file | `--output=<file>` |
| `--align` | Align the output | `--align=<type>` |

## CLI Examples

Basic usage:

```bash
go run ./cmd/ascii-art "hello"
```

With banner selection:

```bash
go run ./cmd/ascii-art "hello" shadow
```

With color (whole string):

```bash
go run ./cmd/ascii-art --color=red "hello"
```

With color (substring mode):

```bash
go run ./cmd/ascii-art --color=green "ll" "hello"
```

With alignment:

```bash
go run ./cmd/ascii-art --align=center "hello"
```

With file output:

```bash
go run ./cmd/ascii-art --output=result.txt "hello"
```

Multiline example:

```bash
go run ./cmd/ascii-art "Hello\nThere"
```

Special characters example:

```bash
go run ./cmd/ascii-art '!@#$%^&*()'
```

Using `make`:

```bash
make run ARGS="hello"
```

## Web Interface

Start the web server:

```bash
go run ./cmd/ascii-art-web
```

Then open:

```text
http://localhost:8080
```

The web interface supports:
- text input
- banner selection (`standard`, `shadow`, `thinkertoy`)
- color selection through the color wheel UI
- live rendering mode (instant results via API as you type)
- manual mode (Generate button submits the form)
- copy result to clipboard
- export result as a downloadable `.txt` file

### Web Routes

- `GET /`: serves the main HTML page
- `POST /ascii-art`: processes the form and renders the result in HTML
- `POST /api/ascii-art`: accepts JSON and returns JSON
- `POST /export`: renders ASCII art and returns it as a downloadable `ascii_art.txt` file
- `/assets/`: serves static assets

> Note: alignment is supported in the CLI only. The web interface does not expose alignment controls.

## API Usage

The project exposes a JSON API endpoint:

```text
POST /api/ascii-art
```

### Example Request

```bash
curl -X POST http://localhost:8080/api/ascii-art \
  -H "Content-Type: application/json" \
  -d '{
    "text": "hello",
    "banner": "standard",
    "align": "left",
    "color": "#ff6600",
    "color_substring": ""
  }'
```

### Example Response

```json
{
  "result": "ASCII OUTPUT HERE"
}
```

### API Validation

- `text` and `banner` are required
- `banner` must be one of `standard`, `shadow`, `thinkertoy`
- `align` is optional; must be one of `left`, `center`, `right`, `justify` if provided
- input must stay within ASCII `32..126` plus newline support
- malformed JSON returns `400 Bad Request`

### HTTP Status Codes

- `200 OK`: successful request
- `400 Bad Request`: invalid input, malformed JSON, invalid banner/color/align
- `404 Not Found`: unknown path or missing banner file
- `405 Method Not Allowed`: unsupported HTTP method
- `500 Internal Server Error`: template, banner parsing, or render failure

## Docker

Build the Docker image:

```bash
docker build -t ascii-art-web:latest .
```

Run the container:

```bash
docker run --rm -d -p 8080:8080 --name ascii-art-web ascii-art-web:latest
```

Then open:

```text
http://localhost:8080
```

Verify the API from your host:

```bash
curl -s -X POST http://localhost:8080/api/ascii-art \
  -H "Content-Type: application/json" \
  -d '{"text":"Docker","banner":"standard"}' \
  | python3 -c "import sys,json; print(json.load(sys.stdin)['result'])"
```

Stop the container (if started without `--rm`):

```bash
docker stop ascii-art-web
```

Clean up unused Docker objects:

```bash
docker image prune -f
docker container prune -f
```

Run the helper script:

```bash
./dockerize.sh
```

Audit-style verification commands:

```bash
docker image build -f Dockerfile -t ascii-art-web-docker .
docker images
docker container run -d -p 8080:8080 --name dockerize ascii-art-web-docker
docker ps -a
docker exec -it dockerize /bin/bash
cd /app
ls -l
```

Expected runtime layout inside the container:

- `/app/server`
- `/app/templates`
- `/app/assets`

Docker notes:

- The runtime image uses only the compiled web binary plus `templates/` and `assets/`.
- The container runs as the non-root user `appuser`.
- The image includes Docker labels for `maintainer`, `version`, and `description`.

## Testing

Run all tests:

```bash
go test ./...
```

This includes:
- unit tests for `internal/banner`, `internal/input`, `internal/output`, and `internal/render`
- handler tests for `internal/web`
- golden integration tests in `test/integration_test.go`

## Error Cases

The CLI exits with code `1` and prints an error when:

- input arguments are missing or invalid
- input contains non-ASCII characters (outside `32..126`, excluding newline)
- banner name is unknown or banner file cannot be read/is malformed
- color format is invalid or unrecognized

The web server and API return HTTP errors for invalid input, missing assets, or render failures.

## Notes

- Default banner is `standard` (`assets/banners/standard.txt`)
- Terminal width for CLI alignment is detected via `tput cols` or `COLUMNS` with fallback `80`
- The web UI uses browser CSS color styling for rendered output
- If rendering logic changes intentionally, regenerate golden files and rerun `go test ./...`
- In interactive `bash`, prefer single quotes when input contains `!` to avoid history expansion issues
