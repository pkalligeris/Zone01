# TASK-03 — Parser: File Reading & Format Validation

## Description

Implement the file-reading and format-validation layer of the `internal/parser` package. This task covers opening the file, reading raw content, and enforcing structural rules (correct characters, 4×4 dimensions, one blank line between pieces, no trailing blank lines).

## Scope

Implement in `internal/parser/parser.go`:

- `Parse(filePath string) ([]Tetromino, error)` — the main entry point for the package
- File opening and full content reading using `os.Open` + `bufio.Scanner`
- Block splitting: split content on single blank lines (`\n\n`), reject double blank lines or trailing blank lines
- Per-block validation:
  - Exactly 4 rows per block
  - Exactly 4 columns per row
  - Only `.` and `#` characters allowed (reject everything else including spaces)
- Return an `error` (not panic) on any violation

> **Note:** Block count validation (exactly 4 `#`) and contiguity checking are covered in TASK-04.

## Dependencies

- **TASK-02** (the `Tetromino` type must exist)

## Acceptance Criteria

- [ ] `Parse` opens the file at the given path; returns an error if the file does not exist or cannot be read
- [ ] Returns an error if the file contains zero blocks
- [ ] Returns an error if any character is not `.`, `#`, or `\n`
- [ ] Returns an error if any block has ≠ 4 rows
- [ ] Returns an error if any row has ≠ 4 columns
- [ ] Returns an error if two consecutive blank lines appear between blocks
- [ ] Returns an error if a trailing blank line appears at end of file
- [ ] All valid fixture files in `testdata/` parse without error

## TDD — Write Tests First

Write the following tests in `parser_test.go` **before** implementing `Parse`:

| Test name | Input | Expected |
|---|---|---|
| `TestParse_FileNotFound` | `"nonexistent.txt"` | `error != nil` |
| `TestParse_EmptyFile` | `testdata/` file with 0 blocks | `error != nil` |
| `TestParse_InvalidChar` | `testdata/invalid_chars.txt` | `error != nil` |
| `TestParse_InvalidDimensions` | `testdata/invalid_dimensions.txt` | `error != nil` |
| `TestParse_InvalidSpacing` | `testdata/invalid_spacing.txt` | `error != nil` |
| `TestParse_ValidSingle` | `testdata/valid_single.txt` | `len(result) == 1, err == nil` |
| `TestParse_ValidMulti` | `testdata/valid_multi.txt` | `len(result) > 1, err == nil` |

## Definition of Done

- All tests above pass with `go test -v ./internal/parser/`
- No test relies on block count or contiguity (those are TASK-04)
- Uses only standard library: `os`, `bufio`, `strings`, `fmt`

## Reference

- [Architecture.md — §8.1 TDD Test Scenarios (parser)](../Architecture.md)
- [PRD.md — §3.1 Input Processing](../PRD.md)
- [PRD.md — §4 Error Handling](../PRD.md)
