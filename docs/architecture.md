# System Architecture

## Overview
The ASCII Art Generator follows the **Standard Go Project Layout**, ensuring scalability, maintainability, and clear separation of concerns. The project provides two entrypoints: a **CLI application** (`cmd/ascii-art`) and a **Web server** (`cmd/ascii-art-web`). Both share the same core packages for banner loading, rendering, and data models.

## Container Runtime

The Docker image is a multi-stage build. The builder stage compiles the Go web binary, while the runtime stage contains only the files required to serve the application during evaluation:

- `/app/server` - compiled web server binary
- `/app/templates/` - HTML templates
- `/app/assets/` - static assets and banner files

The final image runs as the non-root user `appuser`, exposes port `8080`, and includes `/bin/bash` so an auditor can inspect the filesystem with `docker exec`.

## Component Diagram

```mermaid
graph TD
    CLI[cmd/ascii-art] -->|Calls| Input[internal/input]
    CLI -->|Calls| Banner[internal/banner]
    CLI -->|Calls| Render[internal/render]
    
    Input -->|Returns Config Object| CLI
    Banner -->|Returns Banner Map| CLI
    
    CLI -->|Config + Banner| Render
    Render -->|Returns Art| CLI
    CLI -->|Stdout or File| User

    Web[cmd/ascii-art-web] -->|Imports| InternalWeb[internal/web]
    InternalWeb -->|Calls| Banner
    InternalWeb -->|Calls| Render
    InternalWeb -->|Renders| Templates[templates/index.html]
    InternalWeb -->|Serves| Assets[assets/]
    InternalWeb -->|HTTP Response| Browser[Browser]
    InternalWeb -->|File Download| Browser[Browser]
    InternalWeb -->|JSON API Response| APIClient[API Client]
    
    Test[test/integration] -->|Validates| CLI
```

## Modules

### 1. CLI Controller (`cmd/ascii-art/main.go`)
*   **Responsibility:** CLI orchestration.
*   **Flow:**
    1.  Calls `Input Parser` to validate and sanitize arguments.
    2.  Calls `Banner Loader` to read the font file.
    3.  Passes the string and banner to the `Renderer`.
    4.  Prints the result to the console.

### 2. Web Layer (`internal/web` & `cmd/ascii-art-web/main.go`)
*   **Responsibility:** HTTP server orchestration and RESTful JSON API.
*   **Components:**
    *   `cmd/ascii-art-web/main.go` — Entrypoint. Wires up routes and starts the ":8080" server.
    *   `internal/web/handlers.go` — HTTP handlers (`GET /`, `POST /ascii-art`, `POST /api/ascii-art`, `POST /export`).
    *   `internal/web/service.go` — Web-specific core logic (parsing form data, calling the renderer, managing App errors).
    *   Static file server — serves `/assets/` for CSS, background images, etc.
*   **Flow:**
    1.  Receives HTTP request (Form POST, JSON API, or File Export).
    2.  Validates form/JSON input (text, banner name, ASCII limits, 32KB payload caps).
    3.  Calls `Banner Loader` to read the font file.
    4.  Passes the text and banner to the `Renderer`.
    5.  Returns an HTML template, a JSON object, or a downloadable `.txt` file via headers.

### 3. Input Layer (`internal/input`)
*   **Responsibility:** Validation and Sanitization (CLI only).
*   **Key Functions:**
    *   `ParseArgs(args []string) (*Config, error)`: Parses flags (`--color`, `--output`, `--align`), handles positional arguments, and validates input.

### 4. Data Layer (`internal/banner`)
*   **Responsibility:** Data Access.
*   **Key Functions:**
    *   `LoadBanner(filename string) (Banner, error)`: Reads the file, splits by double newline (or specific separator), and maps runes to 8-line string slices.

### 5. Logic Layer (`internal/render`)
*   **Responsibility:** Core Logic.
*   **Key Functions:**
    *   `Render(config *Config, b Banner) string`: Generates the ASCII art, applies ANSI color codes if requested, and aligns text based on terminal width.

### 6. Templates (`templates/`)
*   **Responsibility:** HTML presentation for the web interface.
*   **Files:**
    *   `index.html` — Dark-themed form with text input, banner radio buttons, result and error display areas. Rendered via Go `html/template`.

### 7. Shared Models (`pkg/model`)
*   **Responsibility:** Data Contracts.
*   **Data Structure:**
    *   `type Banner map[rune][]string`
    *   `type Config struct { ... }`: Holds Input string, Banner name, Color settings, Output file path, and Alignment mode.

## Directory Structure
```
.
├── cmd/
│   ├── ascii-art/
│   │   └── main.go           # CLI entry point (wiring only)
│   └── ascii-art-web/
│       └── main.go           # Web server entry point
├── internal/                 # Private application code (not importable by others)
│   ├── banner/               # Banner loading and parsing logic
│   │   ├── loader.go
│   │   └── loader_test.go
│   ├── input/                # Input validation and sanitization (CLI)
│   │   ├── parser.go
│   │   └── parser_test.go
│   ├── render/               # ASCII art generation logic
│   │   ├── renderer.go
│   │   └── renderer_test.go
│   └── web/                  # Web server HTTP logic and handlers
│       ├── handlers.go
│       ├── handlers_test.go
│       ├── service.go
│       ├── template.go
│       └── types.go
├── pkg/                      # Public library code (safe for external use)
│   └── model/                # Shared data structures (e.g., Banner type)
├── templates/                # HTML templates for web interface
│   └── index.html
├── test/                     # External integration tests
│   ├── golden/               # Golden file test cases (.txt files)
│   └── integration_test.go   # "The Tester" - runs binary against golden files
├── docs/                     # Project documentation
│   ├── architecture.md       # This file
│   ├── prd.md                # Product Requirements Document
│   ├── tasks/                # Task breakdown (task_01, task_02, etc.)
│   └── logs/                 # AI interaction logs (ai.log)
├── assets/                   # Static assets
│   └── banners/              # standard.txt, shadow.txt, thinkertoy.txt
├── AGENTS.md                 # AI agent instructions (logging, project context)
├── go.mod
└── Makefile                  # Build and test automation
```
