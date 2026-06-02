# TASK-09 — Stress Test & Final Verification

## Description

Validate the complete system end-to-end with the 26-piece stress test, verify all test suites pass with the race detector, and confirm the binary meets all PRD acceptance criteria. This is the final quality gate before the project is considered complete.

## Scope

- Run the full test suite with the race detector: `go test -v -race ./...`
- Run the 26-piece stress test and measure execution time: `time ./tetris-optimizer testdata/valid_26.txt`
- Manually verify the output of `valid_single.txt`, `valid_multi.txt`, and `valid_26.txt` against expected results
- Verify all error paths produce exactly `ERROR\n` (no extra whitespace, no stack traces)
- Verify that the binary exits with code `1` on every error path and `0` on success

## Dependencies

- **TASK-08** (all packages complete and binary buildable)

## Acceptance Criteria

- [ ] `go test -v -race ./...` — zero failures, zero race conditions
- [ ] `time ./tetris-optimizer testdata/valid_26.txt` completes in **< 10 seconds**
- [ ] Output for `valid_26.txt` contains exactly `26` distinct uppercase letters (`A–Z`)
- [ ] Output for `valid_single.txt` is a valid square with exactly one letter repeated
- [ ] Every row in every output ends with `\n`
- [ ] All 8 error fixture files (`invalid_*.txt`) produce exactly `ERROR\n` and exit `1`
- [ ] No argument / two arguments → `ERROR\n`, exit `1`
- [ ] `go vet ./...` produces zero warnings
- [ ] `go build ./...` produces no warnings

## Checklist

| Check | Command | Pass condition |
|---|---|---|
| Unit tests | `go test -v ./internal/...` | All pass |
| Integration tests | `go test -v ./cmd/...` | All pass |
| Race detector | `go test -race ./...` | No races detected |
| Vet | `go vet ./...` | Zero output |
| Build | `go build ./...` | No errors |
| Stress test timing | `time ./tetris-optimizer testdata/valid_26.txt` | < 10 s |
| Error output format | Manual / test | Exactly `ERROR\n` |
| Exit codes | Integration test assertions | `0` on success, `1` on error |

## Definition of Done

All checklist items above are green. The project is complete.

## Reference

- [Architecture.md — §10 Build & Test Commands](../Architecture.md)
- [PRD.md — §6 Testing Strategy](../PRD.md)
