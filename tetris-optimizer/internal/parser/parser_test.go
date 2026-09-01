package parser

import (
	"testing"
)

// ── TASK-02: Core data type tests ─────────────────────────────────────────────

// TestTetromino_CellsLength verifies the Tetromino struct has exactly 4 cells.
func TestTetromino_CellsLength(t *testing.T) {
	var piece Tetromino
	if len(piece.Cells) != 4 {
		t.Errorf("expected Tetromino.Cells length 4, got %d", len(piece.Cells))
	}
}

// TestTetromino_CellsAreRowCol verifies each cell is a [2]int (row, col).
func TestTetromino_CellsAreRowCol(t *testing.T) {
	piece := Tetromino{
		Cells: [4][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	}
	for i, cell := range piece.Cells {
		if len(cell) != 2 {
			t.Errorf("cell %d: expected 2 elements, got %d", i, len(cell))
		}
	}
}

// ── TASK-03: File reading & format validation tests ───────────────────────────

// TestParse_FileNotFound returns an error for a missing file.
func TestParse_FileNotFound(t *testing.T) {
	_, err := Parse("nonexistent_file_xyz.txt")
	if err == nil {
		t.Error("expected error for nonexistent file, got nil")
	}
}

// TestParse_EmptyFile returns an error when the file contains no blocks.
func TestParse_EmptyFile(t *testing.T) {
	_, err := Parse("../../testdata/empty.txt")
	if err == nil {
		t.Error("expected error for empty file, got nil")
	}
}

// TestParse_InvalidChar returns an error when a file contains illegal characters.
func TestParse_InvalidChar(t *testing.T) {
	_, err := Parse("../../testdata/invalid_chars.txt")
	if err == nil {
		t.Error("expected error for invalid characters, got nil")
	}
}

// TestParse_InvalidDimensions returns an error when a block is not 4 rows.
func TestParse_InvalidDimensions(t *testing.T) {
	_, err := Parse("../../testdata/invalid_dimensions.txt")
	if err == nil {
		t.Error("expected error for invalid dimensions, got nil")
	}
}

// TestParse_InvalidSpacing returns an error for double blank lines or trailing blank lines.
func TestParse_InvalidSpacing(t *testing.T) {
	_, err := Parse("../../testdata/invalid_spacing.txt")
	if err == nil {
		t.Error("expected error for invalid spacing, got nil")
	}
}

// TestParse_ValidSingle parses a single valid tetromino without error.
func TestParse_ValidSingle(t *testing.T) {
	pieces, err := Parse("../../testdata/valid_single.txt")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if len(pieces) != 1 {
		t.Errorf("expected 1 piece, got %d", len(pieces))
	}
}

// TestParse_ValidMulti parses multiple valid tetrominoes without error.
func TestParse_ValidMulti(t *testing.T) {
	pieces, err := Parse("../../testdata/valid_multi.txt")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if len(pieces) < 2 {
		t.Errorf("expected at least 2 pieces, got %d", len(pieces))
	}
}

// ── TASK-04: Shape validation tests ──────────────────────────────────────────

// TestParse_InvalidBlockCount_Low returns an error for a grid with 3 hashes.
func TestParse_InvalidBlockCount_Low(t *testing.T) {
	_, err := Parse("../../testdata/invalid_blocks.txt")
	if err == nil {
		t.Error("expected error for block with 3 '#' cells, got nil")
	}
}

// TestParse_InvalidBlockCount_High returns an error for a grid with 5 hashes.
func TestParse_InvalidBlockCount_High(t *testing.T) {
	_, err := Parse("../../testdata/invalid_blocks_high.txt")
	if err == nil {
		t.Error("expected error for block with 5 '#' cells, got nil")
	}
}

// TestParse_Disjoint returns an error when the '#' cells are not contiguous.
func TestParse_Disjoint(t *testing.T) {
	_, err := Parse("../../testdata/invalid_disjoint.txt")
	if err == nil {
		t.Error("expected error for disjoint '#' cells, got nil")
	}
}

// TestParse_TooManyPieces returns an error when the file has 27 tetrominoes.
func TestParse_TooManyPieces(t *testing.T) {
	_, err := Parse("../../testdata/invalid_too_many.txt")
	if err == nil {
		t.Error("expected error for 27 tetrominoes, got nil")
	}
}

// TestParse_IPiece_Normalised verifies the I-piece is normalised to (0,0)..(0,3).
func TestParse_IPiece_Normalised(t *testing.T) {
	pieces, err := Parse("../../testdata/valid_single.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := [4][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	if pieces[0].Cells != expected {
		t.Errorf("I-piece cells: got %v, want %v", pieces[0].Cells, expected)
	}
}

// TestParse_LPiece_Normalised verifies the L-piece top-left cell is (0,0).
func TestParse_LPiece_Normalised(t *testing.T) {
	pieces, err := Parse("../../testdata/valid_l_piece.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Top-left occupied cell must be (0,0).
	minR, minC := pieces[0].Cells[0][0], pieces[0].Cells[0][1]
	for _, cell := range pieces[0].Cells {
		if cell[0] < minR {
			minR = cell[0]
		}
		if cell[1] < minC {
			minC = cell[1]
		}
	}
	if minR != 0 || minC != 0 {
		t.Errorf("L-piece not normalised: top-left is (%d,%d), want (0,0)", minR, minC)
	}
}

// TestParse_26Pieces successfully parses exactly 26 tetrominoes.
func TestParse_26Pieces(t *testing.T) {
	pieces, err := Parse("../../testdata/valid_26.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(pieces) != 26 {
		t.Errorf("expected 26 pieces, got %d", len(pieces))
	}
}
