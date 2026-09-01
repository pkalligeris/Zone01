# Task 01: Banner Management (Developer 1)

**Objective:** Implement the Data Layer (`internal/banner`) responsible for reading `standard.txt` and parsing it into the `model.Banner` map.

## TDD Cycle (Red-Green-Refactor)

### Cycle 1: File Loading
1.  **RED (Write Test):**
    *   Create `internal/banner/loader_test.go`.
    *   Write `TestLoadBanner_MissingFile`. Call `LoadBanner("nonexistent.txt")` and assert that it returns an error.
2.  **GREEN (Write Code):**
    *   In `internal/banner/loader.go`, implement `LoadBanner` using `os.ReadFile`.
    *   Return an error if the read fails.
3.  **REFACTOR:**
    *   Ensure error messages are clear.

### Cycle 2: Parsing Logic
1.  **RED (Write Test):**
    *   Write `TestLoadBanner_ValidParsing`.
    *   Create a temporary test file (or use a mock string) containing a single character in the standard format (8 lines of art).
    *   Call `LoadBanner` on this file.
    *   Assert that the returned `Banner` map contains the expected key (e.g., ' ') and that the value is a slice of 8 strings.
2.  **GREEN (Write Code):**
    *   Implement the parsing logic in `internal/banner/loader.go`.
    *   **Logic:**
        *   Split the file content by newlines (`\n`).
        *   Iterate through the lines.
        *   The file structure is usually: 8 lines of character data, followed by 1 empty line separator.
        *   Map the chunks to ASCII characters starting from 32 (Space) to 126 (~).
3.  **REFACTOR:**
    *   Handle potential Windows line endings (`\r\n`) by sanitizing the input.
    *   Ensure the loop handles the end of the file correctly.

## Verification
*   Run `go test ./internal/banner/... -v` to confirm all tests pass.
*   Ensure `LoadBanner` returns `pkg/model.Banner`.

## Acceptance Criteria
*   [x] `LoadBanner` correctly reads a file from disk.
*   [x] The returned `Banner` map contains entries for ASCII 32 through 126.
*   [x] Each map value (slice of strings) has exactly 8 lines.
*   [x] Returns a specific error if the file does not exist or is empty.
*   [x] Unit tests cover valid parsing and error conditions.
