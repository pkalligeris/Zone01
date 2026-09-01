# TASK-08 — Main: CLI Entry Point & Orchestration

## Description

Implement `cmd/tetris-optimizer/main.go` — the entry point that ties all packages together. It validates CLI arguments, delegates to parser → solver → printer, and outputs `ERROR\n` to stdout on any failure, exiting with a non-zero status code.

## Scope

Implement in `cmd/tetris-optimizer/main.go`:

```
os.Args check
    └── parser.Parse(filePath)
            └── solver.Solve(pieces)
                    └── printer.Print(grid)
```

- Validate `len(os.Args) == 2`; print `ERROR` and exit `1` otherwise
- Call `parser.Parse(os.Args[1])`; print `ERROR` and exit `1` on error
- Call `solver.Solve(pieces)`; print `ERROR` and exit `1` on error
- Call `printer.Print(grid)` on success
- `ERROR` must be printed to `os.Stdout` followed by `\n`
- Exit code must be non-zero (`os.Exit(1)`) on any error path

## Dependencies

- **TASK-04** (parser fully implemented)
- **TASK-06** (solver fully implemented)
- **TASK-07** (printer fully implemented)

## Acceptance Criteria

- [ ] Running with no arguments prints `ERROR\n` and exits with code `1`
- [ ] Running with two arguments prints `ERROR\n` and exits with code `1`
- [ ] Running with a non-existent file path prints `ERROR\n` and exits with code `1`
- [ ] Running with `testdata/valid_single.txt` prints a valid square and exits `0`
- [ ] Running with `testdata/valid_multi.txt` prints a valid square and exits `0`
- [ ] Running with any invalid fixture file prints `ERROR\n` and exits `1`
- [ ] Binary compiles with `go build -o tetris-optimizer ./cmd/tetris-optimizer/`

## TDD — Write Tests First

Write an integration test file `cmd/tetris-optimizer/main_test.go`:

| Test name | Command args | Expected stdout | Expected exit |
|---|---|---|---|
| `TestMain_NoArgs` | `[]` | `"ERROR\n"` | `1` |
| `TestMain_TwoArgs` | `["a.txt", "b.txt"]` | `"ERROR\n"` | `1` |
| `TestMain_FileNotFound` | `["missing.txt"]` | `"ERROR\n"` | `1` |
| `TestMain_InvalidChars` | `["testdata/invalid_chars.txt"]` | `"ERROR\n"` | `1` |
| `TestMain_ValidSingle` | `["testdata/valid_single.txt"]` | valid grid output | `0` |
| `TestMain_ValidMulti` | `["testdata/valid_multi.txt"]` | valid grid output | `0` |

> Use `os/exec` to run the compiled binary in integration tests for true end-to-end coverage.

## Definition of Done

- `go build -o tetris-optimizer ./cmd/tetris-optimizer/` succeeds
- All integration tests pass with `go test -v ./cmd/tetris-optimizer/`
- All prior unit tests still pass (no regression): `go test ./...`
- `go test -race ./...` passes cleanly

## Reference

- [Architecture.md — §5 Data Flow](../Architecture.md)
- [Architecture.md — §7 Error Handling Contract](../Architecture.md)
- [Architecture.md — §8.4 Integration TDD Test Scenarios](../Architecture.md)
- [PRD.md — §4 Error Handling](../PRD.md)
