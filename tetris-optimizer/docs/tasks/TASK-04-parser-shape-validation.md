# TASK-04 — Parser: Shape Validation (Block Count & Contiguity)

## Description

Implement the semantic validation layer of the `internal/parser` package. After structural parsing (TASK-03), each `4×4` block must be verified to contain exactly 4 `#` characters that are connected (contiguous). Additionally enforce the 26-piece maximum limit.

## Scope

Extend `internal/parser/parser.go`:

- `countBlocks(grid [4]string) int` — counts `#` characters in a 4×4 block
- `isContiguous(grid [4]string) bool` — BFS/DFS flood-fill from the first `#` to verify all 4 are reachable from each other
- `normalise(grid [4]string) [4][2]int` — extracts `(row, col)` positions of each `#`, shifts them so the top-left occupied cell is `(0,0)`, returns sorted `[4][2]int`
- Integrate these into `Parse`: after structural checks, apply count → contiguity → normalise per block
- After parsing all blocks, enforce `len(tetrominoes) <= 26`; return an error if exceeded

## Dependencies

- **TASK-03** (structural parsing must be complete)

## Acceptance Criteria

- [ ] Returns an error if any block has fewer than 4 `#` characters
- [ ] Returns an error if any block has more than 4 `#` characters
- [ ] Returns an error if the 4 `#` characters in any block are not contiguous (BFS/DFS must confirm connectivity)
- [ ] Returns an error if more than 26 tetrominoes are present in the file
- [ ] Valid shapes are correctly normalised — top-left `#` maps to `(0,0)`
- [ ] The I-piece (`####` on row 0) normalises to `[0,0],[0,1],[0,2],[0,3]`
- [ ] The returned `[]Tetromino` slice preserves input order

## TDD — Write Tests First

Write the following tests in `parser_test.go` **before** implementing the validation functions:

| Test name | Input | Expected |
|---|---|---|
| `TestParse_InvalidBlockCount_Low` | `testdata/invalid_blocks.txt` (3 hashes) | `error != nil` |
| `TestParse_InvalidBlockCount_High` | Block with 5 hashes | `error != nil` |
| `TestParse_Disjoint` | `testdata/invalid_disjoint.txt` | `error != nil` |
| `TestParse_TooManyPieces` | File with 27 valid pieces | `error != nil` |
| `TestParse_IPiece_Normalised` | Single horizontal I-piece | `Cells == [0,0],[0,1],[0,2],[0,3]` |
| `TestParse_LPiece_Normalised` | L-shaped piece | Top-left cell is `(0,0)` |
| `TestParse_26Pieces` | `testdata/valid_26.txt` | `len(result) == 26, err == nil` |

## Definition of Done

- All tests above pass with `go test -v ./internal/parser/`
- All TASK-03 tests still pass (no regression)
- `go test -race ./internal/parser/` passes cleanly

## Reference

- [Architecture.md — §8.1 TDD Test Scenarios (parser)](../Architecture.md)
- [PRD.md — §3.2 Core Logic — Shape Validation](../PRD.md)
- [PRD.md — §5.2 Maximum Limits](../PRD.md)
