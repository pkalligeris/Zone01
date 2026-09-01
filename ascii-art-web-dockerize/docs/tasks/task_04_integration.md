# Task 04: Integration & Golden Tests (Tester / Joint)

**Objective:** Implement the "Golden File" testing strategy to ensure the application output matches expected results exactly. This connects all modules via `cmd/ascii-art/main.go` and tests the binary.

## Steps

### 1. Wire the Application (`cmd/ascii-art/main.go`)
*   **Action:** Implement `main` to call `input.ParseInput`, `banner.LoadBanner` (using `assets/banners/standard.txt`), and `render.Render`.
*   **Verification:** Run `go run ./cmd/ascii-art "hello"` and visually check output.

### 2. Generate Golden Files
*   **Action:** Create the directory `test/golden/`.
*   **Action:** For each case in `golden_tests.md` (GT-01 to GT-10), run the program and save the output.
    *   Example: `go run ./cmd/ascii-art "hello" > test/golden/hello.txt`
*   **Verification:** Manually open each `.txt` file in `test/golden/` and verify the ASCII art looks correct.

### 3. Implement Test Runner (`test/integration_test.go`)
*   **RED (Write Test):**
    *   Create `test/integration_test.go`.
    *   Define a struct for test cases: `struct { Name, Input, GoldenFile }`.
    *   Create a slice of these structs matching `golden_tests.md`.
    *   In `TestGolden`, loop through cases:
        1.  Run the main program (using `exec.Command` or by calling a refactored `Run()` function in main).
        2.  Capture `stdout`.
        3.  Read the expected file from `test/golden/`.
        4.  Compare actual vs expected.
*   **GREEN (Fix Bugs):**
    *   If tests fail, debug `main.go` or specific modules.
    *   Ensure newlines at the end of the output match exactly.

### 4. Continuous Integration Check
*   **Action:** Run the full suite.
    *   `go test ./...` (Runs unit tests in `internal/` and integration tests in `test/`).

## Deliverables
*   `cmd/ascii-art/main.go` (Working entry point).
*   `test/golden/*.txt` (Verified reference files).
*   `test/integration_test.go` (Automated regression suite).

## Acceptance Criteria
*   [x] `go run ./cmd/ascii-art "hello"` prints the ASCII art to stdout.
*   [x] Golden files (`test/golden/*.txt`) exist for all 10 test cases.
*   [x] `go test ./test/...` passes, verifying actual output matches golden files.
*   [x] The application defaults to `assets/banners/standard.txt`.
*   [x] CI check (`go test ./...`) passes for the whole project.
