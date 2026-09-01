# Task 03: Rendering Engine (Developer 3)

**Objective:** Implement the Logic Layer (`internal/render`) responsible for generating the final ASCII art string using the Banner map.

## TDD Cycle (Red-Green-Refactor)

### Cycle 1: Single Line Rendering
1.  **RED (Write Test):**
    *   Create `internal/render/renderer_test.go`.
    *   Write `TestRender_Simple`.
    *   **Mocking:** Manually construct a `Banner` map in the test (do not rely on `LoadBanner`). Add one character (e.g., 'A') with 8 dummy lines.
    *   Call `Render("A", mockBanner)`.
    *   Assert that the output string matches the 8 dummy lines joined by newlines.
2.  **GREEN (Write Code):**
    *   In `internal/render/renderer.go`, implement `Render`.
    *   Initialize a slice of 8 strings (representing the 8 rows of output).
    *   Loop through the input string. For each character, append its 8 lines to the corresponding 8 rows of the output.
    *   Join the 8 rows with `\n` and return.

### Cycle 2: Multi-line Rendering
1.  **RED (Write Test):**
    *   Write `TestRender_MultiLine`.
    *   Input: `"A\nB"`.
    *   Assert that the output contains the block for 'A', followed by a newline, followed by the block for 'B'.
2.  **GREEN (Write Code):**
    *   Update `Render` to split the input string by `\n`.
    *   Process each segment (line of text) separately using the logic from Cycle 1.
    *   Join the resulting blocks.

### Cycle 3: Edge Cases
1.  **RED (Write Test):**
    *   Write `TestRender_EmptyLines`. Input: `"\n"`. Assert it prints a newline.
2.  **GREEN (Write Code):**
    *   Ensure that if an input segment is empty, the renderer prints a blank line instead of nothing.
3.  **REFACTOR:**
    *   Clean up string concatenation (consider using `strings.Builder` for performance).

## Verification
*   Run `go test ./internal/render/... -v` to confirm all tests pass.
*   Ensure `Render` accepts `pkg/model.Banner`.

## Acceptance Criteria
*   [ ] `Render` function accepts a string and a `Banner` map.
*   [x] Output string consists of 8 lines of height for every line of input text.
*   [x] Characters are concatenated horizontally with correct spacing.
*   [x] Newlines in input result in a new 8-line block in output.
*   [x] Empty input strings result in no output (or just a newline if appropriate).
*   [x] Unit tests cover single line, multi-line, and empty inputs.
