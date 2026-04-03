# 📄 PRD – ASCII Art Generator (Go)

## 1. Overview

The ASCII Art Generator is a feature-rich application written in Go. It converts input strings into graphical ASCII-art using predefined banner files, with support for text colorization, output alignment, and file export. The project provides both a **command-line interface** and a **web interface** for generating ASCII art.

The program must render letters, numbers, spaces, special characters, and newline sequences (`\n`) according to the format defined in the provided banner files.

---

## 2. Objectives

- Convert input text into ASCII art format.
- Support banner templates:
  - `standard`
  - `shadow`
  - `thinkertoy`
- **Colorize output** based on user input (whole string or substrings).
- **Export output** to a file.
- **Align text** (left, center, right, justify) relative to terminal size.
- **Web interface** for browser-based ASCII art generation.
- **Containerise the application** using Docker for portable, reproducible deployments.
- Ensure output strictly matches expected formatting examples.
- Maintain clean, modular, and testable Go code.

---

## 3. Scope

### In Scope

- Parsing CLI input.
- Parsing flags (`--color`, `--output`, `--align`).
- Interpreting literal `\n` as line breaks.
- Loading banner files from the filesystem.
- Rendering ASCII characters (ASCII 32–126).
- Applying ANSI color codes to output.
- Writing output to files.
- Detecting terminal window size for alignment.
- Web interface with HTTP server, HTML form, and template rendering.
- Supporting:
  - Uppercase letters
  - Lowercase letters
  - Numbers
  - Spaces
  - Special characters
- Handling multiple consecutive newline sequences.
- Unit testing core logic.
- Containerising the application with a multi-stage `Dockerfile`.
- Running the web server as a non-root user inside a minimal Alpine runtime image.

### Out of Scope

- Generating ASCII art algorithmically.
- Editing or modifying banner files.
- Supporting characters outside ASCII 32–126.

---

## 4. Functional Requirements

### 4.1 General Input Handling

- The program must accept flags and positional arguments.
- The input may contain:
  - Letters
  - Numbers
  - Spaces
  - Special characters
  - Literal `\n` sequences
- The program must:
  - Interpret `\n` as a new line
  - Handle multiple consecutive `\n`
  - Handle empty string input

### 4.2 Color Feature

- **Flag:** `--color=<color>`
- **Syntax:** `go run . --color=<color> <substring> "string"`
- **Behavior:**
  - Colors the specified `<substring>` within the string.
  - If `<substring>` is not provided (or matches the main string), the whole string is colored.
  - Supports standard color names (red, blue, etc.), Hex (#RRGGBB), RGB (rgb(r,g,b)), and HSL (hsl(h,s,l)).
  - **Ambiguity Handling:** When 2 arguments are provided (`go run . --color=red arg1 arg2`), the program checks if `arg2` is a valid banner.
    - If valid banner: `Input=arg1`, `Banner=arg2`.
    - If not: `ColorSubstr=arg1`, `Input=arg2`.
 - **Usage Error:** If the flag format is incorrect or color is invalid:
  ```text
  Usage: go run . --color=<color> <substring> "string"

  Supported formats: ANSI standard colors (red, green, blue...), Hex (#RRGGBB), RGB (rgb(r,g,b)), HSL (hsl(h,s,l))
  ```

### 4.3 Output File Feature

- **Flag:** `--output=<fileName.txt>`
- **Syntax:** `go run . --output=<fileName.txt> [STRING] [BANNER]`
- **Behavior:**
  - Writes the resulting ASCII art to `<fileName.txt>` instead of stdout.
  - Supports optional `[BANNER]` argument.
- **Usage Error:** If the flag format is incorrect:
  ```text
  Usage: go run . [OPTION] [STRING] [BANNER]
  
  EX: go run . --output=<fileName.txt> something standard
  ```

### 4.4 Alignment Feature

- **Flag:** `--align=<type>`
- **Syntax:** `go run . --align=<type> [STRING] [BANNER]`
- **Types:**
  - `left` (default behavior)
  - `center`
  - `right`
  - `justify`
- **Behavior:**
  - Adapts the graphical representation to the current terminal size.
  - Only text that fits the terminal size will be tested.
- **Usage Error:** If the flag format is incorrect:
  ```text
  Usage: go run . [OPTION] [STRING] [BANNER]
  
  Example: go run . --align=right something standard
  ```

### 4.5 Banner Selection

Banner files are preformatted ASCII templates.
The user can specify a banner as the last argument.

- **Syntax:** `go run . [STRING] [BANNER]`
- **Defaults:** If not specified, use `standard`.
- **Usage Error:** If the format is incorrect:
  ```text
  Usage: go run . [STRING] [BANNER]
  
  EX: go run . something standard
  ```

---

### 4.6 Rendering Rules

Characters must be rendered horizontally.

Each output block must contain exactly 8 lines per input line.

Rendering must preserve:

- Case sensitivity
- Spacing

Empty input lines must result in blank output lines.

Output must match provided examples exactly.

---

### 4.8 Web Interface

- **Entrypoint:** `cmd/ascii-art-web/main.go` acting as wrapper for `internal/web` package.
- **Port:** HTTP server listening on `:8080`.
- **Routes:**
  - `GET /` — Serves the home page with an ASCII art generation form.
  - `POST /ascii-art` — Accepts form input, generates ASCII art, and returns an HTML result.
  - `POST /api/ascii-art` — JSON API endpoint for generating ASCII art programmatically.
  - `POST /export` — Endpoint returning a downloadable `.txt` file with HTTP headers (`Content-Disposition: attachment`).
  - `/assets/` — Serves static files (CSS, Javascript, background images).
- **Template:** `templates/index.html` rendered via Go `html/template`.
- **Form Fields:**
  - `text` (textarea) — The input string to convert.
  - `banner` (radio buttons) — Banner selection: `standard`, `shadow`, `thinkertoy`.
  - `color` (color picker) — Select hex colors for styling the ASCII output.
- **Input Validation & Security:**
  - Text and banner must not be empty (→ 400 Bad Request).
  - API payloads are strictly hard-capped at 32KB to prevent memory exhaustion DoS attacks.
  - Characters must be within ASCII 32–126 or newline (→ 400 Bad Request).
- **Error Handling:**
  - Banner file not found → 404 Not Found.
  - Failed banner load or render → 500 Internal Server Error.
  - Invalid HTTP method → 405 Method Not Allowed.
- **Behavior:**
  - Features real-time **AJAX Live Preview** updating the display as the user types without full page reloads.
  - Utilizes modern Glassmorphism CSS design.
  - Allows exporting the finalized output structure as a `.txt` file downloading directly to the physical disk.

---

### 4.9 Containerisation (Docker)

- **Dockerfile strategy:** Multi-stage build to keep the final image small and production-ready.
  - **Stage 1 — Builder** (`golang:1.22-alpine`):
    - Sets `WORKDIR /build`.
    - Copies `go.mod` first to leverage Docker layer caching (dependencies only re-downloaded when `go.mod` changes).
    - Copies the full source tree and compiles a statically-linked binary (`CGO_ENABLED=0 GOOS=linux`) targeting `/app/ascii-art-web`.
  - **Stage 2 — Runtime** (`alpine:3.19`):
    - Creates a dedicated non-root system group (`appgroup`) and user (`appuser`) for security hardening.
    - Copies only the compiled binary, `templates/`, and `assets/` from the builder stage.
    - Sets ownership of all copied files to `appuser:appgroup`.
    - Switches to `appuser` before the final `CMD`.
- **Image Metadata Labels** (applied to both stages):
  - `maintainer` — comma-separated list of team members: `nbougial,pkallige,ipapigki`.
  - `version` — current release version (`1.0.0`).
  - `description` — human-readable purpose of the image.
- **Exposed port:** `8080` (matches the web server listen address).
- **Entry point:** `CMD ["./ascii-art-web"]` starts the web server directly.
- **Build command:** `docker build -t ascii-art-web .`
- **Run command:** `docker run --rm -p 8080:8080 ascii-art-web`
- **Workflow:** Each team member builds and runs their own local container from the shared `Dockerfile`; no pre-built image distribution is required.
- **`.dockerignore`:** Excludes files that must not be copied into the build context (e.g. `.git`, local binaries, test artefacts).

---

### 4.7 Error Handling

The program must handle:

- Missing banner file.
- Invalid ASCII characters (outside 32–126).
- Incorrect CLI usage.

Errors must:

- Not crash the program unexpectedly.
- Provide clear, readable error messages.
- Return the specific usage messages defined in sections 4.2, 4.3, 4.4, and 4.5 depending on the context.

### 5. Non-Functional Requirements

- Must use only standard Go packages.
- Code must follow Go best practices.
- Must be modular and maintainable.
- Must include unit tests.
- Must avoid hardcoded ASCII representations inside source code.

---

### 6. Constraints

- Banner height is fixed at 8 lines.
- Characters are separated by newline in banner files.
- ASCII range supported: 32–126.
- Banner files are read-only.
- Implementation must be written in Go.

---

### 7. Acceptance Criteria

The project is considered complete when:

- Output matches all provided examples exactly.
- Multiple newline inputs behave correctly.
- Case sensitivity is preserved.
- **Color flag works for substrings and full strings.**
- **Output flag correctly writes to files.**
- **Alignment flag correctly positions text based on terminal width.**
- **Banner argument correctly switches fonts.**
- **Web interface serves the form, processes input, and supports live preview rendering.**
- **Web interface provides a RESTful JSON API (`/api/ascii-art`).**
- **Web interface allows downloading the generated art as a `.txt` file with correct headers.**
- **Web interface returns correct HTTP status codes for all error conditions.**
- **Docker image builds successfully from the provided `Dockerfile` using `docker build`.**
- **Container starts and serves the web interface on port `8080` via `docker run`.**
- **The process inside the container runs as the non-root `appuser`.**
- **Image metadata labels (`maintainer`, `version`, `description`) are present on the final image.**
- Unit tests pass.
- Code is cleanly modularized following Standard Go Project Layout patterns (`cmd/` and `internal/`).
- No banner data is hardcoded.

---

### 8. Success Metrics

- Zero formatting mismatches in expected output.
- Full support for ASCII 32–126.
- Clear separation of responsibilities in code structure.
- All team members can understand and extend the codebase.
- Docker image size is kept minimal by using a multi-stage build (builder discarded from final image).
- Container runs as a non-root user in production.

---

### 9. Risks & Mitigation

| Risk | Mitigation |
|------|------------|
| Incorrect banner indexing | Write unit tests for specific characters |
| Improper newline handling | Add dedicated newline test cases |
| Merge conflicts between team members | Define API contracts early |
| Output formatting mismatch | Use golden file tests |
| Terminal size detection failure | Fallback to default width (e.g., 80 chars) |
| Docker image bloat from build artefacts | Multi-stage build discards the Go toolchain from the runtime image |
| Container running as root (security risk) | Non-root `appuser` created and enforced via `USER appuser` directive |
| Missing static assets in container | `COPY --from=builder` explicitly copies `templates/` and `assets/` |
| Layer cache invalidation on every build | `go.mod` copied before source so dependency layer is cached independently |

---

### 10. Future Improvements (Optional)

- Add performance optimizations.
- Add additional ASCII fonts.
- Add integration tests.
- Provide export functionality in secondary formats (HTML, JSON).

### 11. Misc.
When Unsure ask for confirmation
Use only standard packages
Do not hallunicate
Stay strict to what is mentioned in the documentation
Do not assume things that are not stated
When updating ai.log keep the format of the file untouched
