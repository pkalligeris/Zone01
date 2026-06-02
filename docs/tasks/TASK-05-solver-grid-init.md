# TASK-05 — Solver: Grid Initialisation & Minimum Size Calculation

## Description

Implement the initialisation logic of the `internal/solver` package: computing the minimum square size and allocating the `Grid`. This is the foundation of the solver before the backtracking algorithm is added in TASK-06.

## Scope

Implement in `internal/solver/solver.go`:

- `MinSize(n int) int` — returns `⌈√(n × 4)⌉` given `n` tetrominoes; uses `math.Sqrt` and `math.Ceil`
- `NewGrid(size int) Grid` — allocates and returns a `size×size` 2D `[][]int` slice filled with `0`
- `Solve(pieces []parser.Tetromino) (Grid, error)` — stub that initialises the grid at `MinSize(len(pieces))` (backtracking logic added in TASK-06)

## Dependencies

- **TASK-02** (the `Tetromino` and `Grid` types must exist)
- **TASK-04** (parser must produce a `[]Tetromino` to pass to the solver)

## Acceptance Criteria

- [ ] `MinSize(1)` returns `2` (`⌈√4⌉ = 2`)
- [ ] `MinSize(4)` returns `4` (`⌈√16⌉ = 4`)
- [ ] `MinSize(5)` returns `5` (`⌈√20⌉ ≈ 4.47 → 5`)
- [ ] `MinSize(26)` returns a reasonable value without panic
- [ ] `NewGrid(3)` returns a `3×3` slice where every cell is `0`
- [ ] `Solve` returns a non-nil `Grid` initialised to `MinSize(len(pieces))` (no placement yet)

## TDD — Write Tests First

Write the following tests in `solver_test.go` **before** implementing:

| Test name | Input | Expected |
|---|---|---|
| `TestMinSize_1` | `n=1` | `2` |
| `TestMinSize_4` | `n=4` | `4` |
| `TestMinSize_5` | `n=5` | `5` |
| `TestMinSize_26` | `n=26` | computed value, no panic |
| `TestNewGrid_Size` | `size=4` | `len(grid)==4`, `len(grid[0])==4` |
| `TestNewGrid_ZeroFilled` | `size=3` | all cells `== 0` |

## Definition of Done

- All tests above pass with `go test -v ./internal/solver/`
- `MinSize` uses only `math` from the standard library
- `NewGrid` makes no assumptions about the backtracking logic

## Reference

- [Architecture.md — §6 Backtracking Algorithm](../Architecture.md)
- [PRD.md — §3.2 Core Logic — Assembly Algorithm](../PRD.md)
