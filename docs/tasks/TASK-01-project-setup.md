# TASK-01 — Project Setup & Scaffolding

## Description

Initialise the Go module and create the full directory/file skeleton for the project as specified in the Architecture. This is the foundation task — all other tasks depend on the structure created here.

## Scope

- Initialise `go.mod` with module name `tetris-optimizer` (no external dependencies)
- Create the following empty placeholder files (with package declarations only):
  - `cmd/tetris-optimizer/main.go`
  - `internal/parser/parser.go` + `parser_test.go`
  - `internal/solver/solver.go` + `solver_test.go`
  - `internal/printer/printer.go` + `printer_test.go`
- Create the `testdata/` directory with all required fixture files:
  - `valid_single.txt` — one valid tetromino
  - `valid_multi.txt` — known multi-piece solvable input
  - `valid_26.txt` — 26-piece stress test input
  - `invalid_chars.txt` — contains an illegal character (e.g. `X`)
  - `invalid_dimensions.txt` — a grid that is not 4×4
  - `invalid_blocks.txt` — a grid with ≠ 4 `#` characters
  - `invalid_disjoint.txt` — blocks that are not contiguous
  - `invalid_spacing.txt` — multiple or trailing blank lines

## Dependencies

None — this is the first task.

## Acceptance Criteria

- [ ] `go mod init tetris-optimizer` runs successfully
- [ ] `go build ./...` compiles without errors (empty stubs are fine)
- [ ] All directories and placeholder files exist at the paths defined in `Architecture.md § 2`
- [ ] All `testdata/` fixture files are present and contain the correct content for their described scenario
- [ ] `go.mod` has zero `require` entries (standard library only)

## Definition of Done

- Project builds cleanly with `go build ./...`
- `go test ./...` runs without compilation errors (tests may fail — that is expected at this stage)
- Fixture files are manually verified to match their described invalid/valid scenario

## Reference

- [Architecture.md — §2 Project Structure](../Architecture.md)
- [PRD.md — §2 Scope & Constraints](../PRD.md)
