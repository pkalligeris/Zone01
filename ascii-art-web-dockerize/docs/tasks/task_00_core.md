# Task 00: Core Infrastructure (Joint Team Task)

**Objective:** Initialize the project and define shared data structures (contracts) so all developers can work independently without blocking each other.

## Steps

1.  **Initialize Module**
    *   Run `go mod init ascii-art`.

2.  **Create Directory Structure**
    *   Create the following folders to match the architecture:
        *   `cmd/ascii-art/`
        *   `internal/banner/`
        *   `internal/input/`
        *   `internal/render/`
        *   `pkg/model/`
        *   `test/golden/`
        *   `assets/banners/`

3.  **Define Shared Types (`pkg/model/banner.go`)**
    *   Create `pkg/model/banner.go`.
    *   Define `package model`.
    *   Define the Banner structure:
        ```go
        // Banner represents the font map. Key is the rune, Value is the 8 lines of ASCII art.
        type Banner map[rune][]string
        ```

4.  **Verification**
    *   Run `go build ./...` to ensure the package structure is valid and there are no syntax errors in the struct definition.

## Acceptance Criteria
*   [x] `go.mod` exists and module name is `ascii-art`.
*   [x] All directories (`cmd`, `internal`, `pkg`, `test`, `assets`) exist.
*   [x] `pkg/model/banner.go` defines `type Banner map[rune][]string`.
*   [x] `go build ./...` runs without errors.
