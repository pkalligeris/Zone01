# TASK-02 — Core Data Types

## Description

Define the shared data types used across all internal packages. These types are the contract between the `parser`, `solver`, and `printer` packages and must be established before any logic is implemented.

## Scope

Define in `internal/parser/parser.go`:

```go
// Tetromino represents a parsed and validated tetris piece.
// Cells holds the (row, col) offsets of each of the 4 blocks,
// normalised so the top-left occupied cell is (0,0).
type Tetromino struct {
    Cells [4][2]int
}
```

Define in `internal/solver/solver.go`:

```go
// Grid is the mutable square board used during backtracking.
// 0 means empty; values 1..N map to tetrominoes A..Z.
type Grid [][]int
```

## Dependencies

- **TASK-01** (project structure must exist)

## Acceptance Criteria

- [ ] `Tetromino` struct is defined in `internal/parser/parser.go` with a `Cells [4][2]int` field
- [ ] `Grid` type is defined in `internal/solver/solver.go` as `[][]int`
- [ ] Both types are exported (uppercase names)
- [ ] `go build ./...` passes with no errors
- [ ] No third-party imports introduced

## TDD — Write Tests First

Before implementing, write the following in `parser_test.go`:

- A test that constructs a `Tetromino` and asserts that `len(t.Cells)` equals `4`
- A test that constructs a `Grid` of size 3×3 and asserts all values are `0`

These tests must **fail** (red) before the types are defined, then **pass** (green) after.

## Definition of Done

- Types compile and are accessible from other internal packages
- Tests in `parser_test.go` and `solver_test.go` pass for type construction
- No behaviour logic added yet (pure data types only)

## Reference

- [Architecture.md — §4 Core Data Types](../Architecture.md)
