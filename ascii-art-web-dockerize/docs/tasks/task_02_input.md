# Task 02: Input Processing (Developer 2)

**Objective:** Implement the Input Layer (`internal/input`) responsible for validating CLI arguments and processing escape sequences.

## TDD Cycle (Red-Green-Refactor)

### Cycle 1: Argument Validation
1.  **RED (Write Test):**
    *   Create `internal/input/parser_test.go`.
    *   Write `TestParseInput_NoArgs`. Pass an empty slice `[]string{}` to `ParseInput`.
    *   Assert that it returns an error (e.g., "usage: go run . [STRING]").
2.  **GREEN (Write Code):**
    *   In `internal/input/parser.go`, implement the check `if len(args) != 1`.

### Cycle 2: Newline Handling
1.  **RED (Write Test):**
    *   Write `TestParseInput_EscapedNewline`.
    *   Pass the string `"Hello\nWorld"` (literal backslash + n).
    *   Assert that the returned string contains the actual newline character (byte 10) or is formatted in a way the renderer expects.
2.  **GREEN (Write Code):**
    *   Use `strings.ReplaceAll(input, "\\n", "\n")` to convert literal escape sequences into actual newlines.

### Cycle 3: ASCII Validation
1.  **RED (Write Test):**
    *   Write `TestParseInput_InvalidChar`.
    *   Pass a string with a non-ASCII character (e.g., "Hell€").
    *   Assert that the function returns an error indicating invalid characters.
2.  **GREEN (Write Code):**
    *   Iterate over the rune slice of the input string.
    *   Check if any rune is outside the range `32` to `126`.
3.  **REFACTOR:**
    *   Optimize the validation loop.

## Verification
*   Run `go test ./internal/input/... -v` to confirm all tests pass.
*   Ensure `ParseInput` correctly handles `\n` vs `\\n`.

## Acceptance Criteria
*   [x] `ParseInput` returns an error if arguments length is not exactly 1.
*   [x] Literal `\n` sequences in input are converted to actual newline characters.
*   [x] Returns an error if input contains characters outside ASCII 32-126.
*   [x] Unit tests cover empty args, valid strings, escaped newlines, and invalid characters.
