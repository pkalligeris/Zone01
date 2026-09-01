# Task 06: Feature Color (Developer 2)

**Objective:** Update the rendering engine to support ANSI color codes.

## TDD Cycle (Red-Green-Refactor)

### Cycle 1: Color Mapping
1.  **RED (Write Test):**
    *   Create `internal/render/color_test.go`.
    *   Write `TestGetColorCode`. Assert "red" returns "\033[31m".
2.  **GREEN (Write Code):**
    *   Create `internal/render/color.go`.
    *   Implement a map or switch statement for standard colors.

### Cycle 2: Color Application
1.  **RED (Write Test):**
    *   Write `TestColorize_Substring`.
    *   Input: "Hello", Substring: "ll", Color: "red".
    *   Assert output has ANSI codes around "ll".
2.  **GREEN (Write Code):**
    *   Implement logic to find indices of substring.
    *   Construct new string with codes inserted.

### Cycle 3: Defaults & Backward Compatibility
1.  **RED (Write Test):**
    *   Write `TestColorize_Empty`.
    *   Input: "Hello", Color: "".
    *   Assert output is identical to input (no codes).
2.  **GREEN (Write Code):**
    *   Add check: if color is empty, return string as-is.
3.  **REFACTOR:**
    *   Integrate into `renderer.go` by updating `Render` signature to accept `Config`.

## Acceptance Criteria
*   [x] Supports standard colors (red, green, blue, etc.).
*   [x] Colors the whole string if no substring is provided.
*   [x] Colors only the substring if provided.
*   [x] Output contains correct ANSI escape sequences.

## Note
*   Coordinate with Developer 4 (Align) as you both modify `renderer.go`.
