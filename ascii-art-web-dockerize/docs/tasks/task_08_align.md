# Task 08: Feature Justify/Align (Developer 4)

**Objective:** Implement text alignment (left, center, right, justify) based on terminal size.

## TDD Cycle (Red-Green-Refactor)

### Cycle 1: Padding Calculation
1.  **RED (Write Test):**
    *   Create `internal/render/align_test.go`.
    *   Write `TestCalculatePadding`.
    *   Mock Terminal Width = 80. Art Width = 20.
    *   Assert Center padding = 30. Right padding = 60.
2.  **GREEN (Write Code):**
    *   Create `internal/render/align.go`.
    *   Implement math for padding.

### Cycle 2: Application
1.  **RED (Write Test):**
    *   Write `TestApplyAlign`.
    *   Input: 8-line ASCII block. Align: "right".
    *   Assert spaces are prepended to each line.
2.  **GREEN (Write Code):**
    *   Implement `ApplyAlign` to iterate lines and add spaces.

### Cycle 3: Defaults
1.  **RED (Write Test):**
    *   Write `TestApplyAlign_Left`.
    *   Assert padding is 0 (string unchanged).
2.  **GREEN (Write Code):**
    *   Handle "left" or empty align by returning input as-is.
3.  **REFACTOR:**
    *   Integrate into `renderer.go`.

## Acceptance Criteria
*   [x] `--align=center` centers the art.
*   [x] `--align=right` aligns to right edge.
*   [x] `--align=justify` spreads words across the line.
*   [x] Adapts to window size (if running in real terminal).

## Note
*   Coordinate with Developer 2 (Color) as you both modify `renderer.go`.
