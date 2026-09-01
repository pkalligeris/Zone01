# Task 07: Feature Output (Developer 3)

**Objective:** Implement file writing capability in the main application controller.

## TDD Cycle (Red-Green-Refactor)

### Cycle 1: File Writer
1.  **RED (Write Test):**
    *   Create `cmd/ascii-art/file_writer_test.go`.
    *   Write `TestWriteOutput`.
    *   Call `WriteOutput("test_output.txt", "content")`.
    *   Assert file exists and contains "content".
2.  **GREEN (Write Code):**
    *   Create `cmd/ascii-art/file_writer.go`.
    *   Implement `WriteOutput` using `os.Create` and `WriteString`.

### Cycle 2: Main Integration
1.  **RED (Write Test):**
    *   Update `test/integration_test.go` (or create a new test case there).
    *   Run binary with `--output=result.txt`.
    *   Assert `result.txt` is created.
2.  **GREEN (Write Code):**
    *   Update `cmd/ascii-art/main.go`.
    *   In `Run`, check `Config.OutputFile`.
    *   If set, call `WriteOutput`. Else, print to stdout.
3.  **REFACTOR:**
    *   Ensure file permissions are correct (0644).
    *   Add a test case in `test/integration_test.go` that runs with `--output` and verifies the file creation and content.

## Acceptance Criteria
*   [x] Flag `--output=<file>` redirects output from stdout to file.
*   [x] File is created if it doesn't exist.
*   [x] File is overwritten if it exists.
*   [x] Permissions are handled gracefully (return error if write fails).

## Note
*   This task depends on Task 05 being complete (or the Config struct being available).