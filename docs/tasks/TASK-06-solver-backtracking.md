# TASK-06 — Solver: Backtracking Algorithm

## Description

Implement the core backtracking placement algorithm in `internal/solver`. Starting from the minimum grid size computed in TASK-05, the solver attempts to place all tetrominoes without overlap. If all positions are exhausted at a given size, the grid grows by one and the solver restarts from piece 0.

## Scope

Implement in `internal/solver/solver.go`:

- `place(pieces []Tetromino, index int, grid Grid) bool`
  - **Base case:** `index == len(pieces)` → return `true`
  - **Find first empty cell** scanning in row-major order (left→right, top→bottom)
  - **For each `(row, col)` offset** of the current piece's `Cells`:
    - Compute the absolute position: `(emptyRow - cellRow, emptyCol - cellCol)`
    - Check all 4 cells fit within bounds and are unoccupied
    - Place the piece (write `index+1` to each cell)
    - Recurse with `place(pieces, index+1, grid)`
    - If recursion returns `false`, undo placement (write `0` back)
  - If no offset fits → return `false`
- Update `Solve(pieces []Tetromino) (Grid, error)`:
  - Start at `size = MinSize(len(pieces))`
  - Loop: call `place`; if it returns `false`, increment `size` and retry with a fresh `NewGrid`
  - Return the solved `Grid` and `nil` error

## Dependencies

- **TASK-05** (grid initialisation and `MinSize` must exist)

## Acceptance Criteria

- [ ] Single I-piece is placed correctly in a `2×2` grid (or next valid size)
- [ ] No two pieces ever share the same cell in the output grid (no overlap)
- [ ] When pieces tile perfectly, no cell in the grid is `0`
- [ ] When pieces cannot tile perfectly, remaining cells are `0` (displayed as `.`)
- [ ] Solver retries with `size+1` if backtracking fails at current size
- [ ] Solves a 5-piece input correctly and matches a known expected output
- [ ] Solves all 26 pieces within 10 seconds (stress)
- [ ] Does not panic on any valid input

## TDD — Write Tests First

Write the following tests in `solver_test.go` **before** implementing backtracking:

| Test name | Scenario | Expected |
|---|---|---|
| `TestSolve_SingleIPiece` | 1 I-piece | Returns solved grid, no error |
| `TestSolve_NoOverlap` | Any valid multi-piece input | No cell value appears more than once per piece index |
| `TestSolve_PerfectFill` | Input that tiles perfectly | Zero cells with value `0` |
| `TestSolve_ImperfectFill` | Input that cannot tile perfectly | Some cells are `0` |
| `TestSolve_SizeGrows` | Input where minimum size is insufficient | `size` of returned grid > `MinSize(n)` |
| `TestSolve_5Pieces` | 5 tetrominoes from `testdata/valid_multi.txt` | Valid solved grid, matches known output |
| `TestSolve_26Pieces` | `testdata/valid_26.txt` | Completes in < 10 s, no panic |

## Definition of Done

- All tests above pass with `go test -v -timeout 30s ./internal/solver/`
- `go test -race ./internal/solver/` passes cleanly
- All TASK-05 tests still pass (no regression)
- Uses only `math` from the standard library (no external packages)

## Reference

- [Architecture.md — §6 Backtracking Algorithm (Mermaid diagram)](../Architecture.md)
- [Architecture.md — §8.2 TDD Test Scenarios (solver)](../Architecture.md)
- [PRD.md — §3.2 Core Logic — Placement Rules](../PRD.md)
