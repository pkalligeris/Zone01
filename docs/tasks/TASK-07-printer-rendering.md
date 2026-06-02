# TASK-07 — Printer: Grid Rendering

## Description

Implement the `internal/printer` package, which converts the solved `Grid` (integer values) into the human-readable output printed to stdout. Each integer `1..N` maps to an uppercase Latin letter (`A..Z`), and `0` maps to `.`.

## Scope

Implement in `internal/printer/printer.go`:

- `Print(grid solver.Grid)` — iterates over every row of the grid:
  - For each cell value `v`:
    - If `v == 0`: write `.`
    - If `v >= 1`: write `string(rune('A' + v - 1))`
  - After each row: write `\n` to stdout
- Output must go to `os.Stdout` using `fmt.Print` or `fmt.Fprintf`

## Dependencies

- **TASK-05** (the `Grid` type must exist in `internal/solver`)

## Acceptance Criteria

- [ ] Cell value `0` renders as `.`
- [ ] Cell value `1` renders as `A`, `2` → `B`, …, `26` → `Z`
- [ ] Every row ends with exactly one `\n`
- [ ] No extra trailing newline after the last row (or consistent with PRD spec)
- [ ] A `2×2` grid with all cells `= 1` produces `AA\nAA\n`
- [ ] A mixed grid renders letters and periods in the correct positions

## TDD — Write Tests First

Write the following tests in `printer_test.go` **before** implementing `Print`. Capture stdout by temporarily redirecting `os.Stdout` to a pipe:

| Test name | Input grid | Expected stdout |
|---|---|---|
| `TestPrint_SinglePiece` | `[[1,1],[1,1]]` | `"AA\nAA\n"` |
| `TestPrint_EmptyCell` | `[[1,0],[0,1]]` | `"A.\n.B\n"` |
| `TestPrint_LetterMapping` | `[[1],[2],[3]]` (3×1) | `"A\nB\nC\n"` |
| `TestPrint_AllPeriods` | `[[0,0],[0,0]]` | `"..\n..\n"` |
| `TestPrint_RowNewlines` | Any grid | Each row ends with `\n` |

## Definition of Done

- All tests above pass with `go test -v ./internal/printer/`
- Output is written to `os.Stdout` only (no `os.Stderr`)
- Uses only `fmt`, `os` from standard library

## Reference

- [Architecture.md — §8.3 TDD Test Scenarios (printer)](../Architecture.md)
- [PRD.md — §3.3 Output Formatting](../PRD.md)
