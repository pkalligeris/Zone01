package printer

import (
	"bytes"
	"io"
	"os"
	"testing"

	"tetris-optimizer/internal/solver"
)

// captureStdout redirects os.Stdout during fn execution and returns what was written.
func captureStdout(fn func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// TestPrint_SinglePiece verifies a 2×2 grid filled with piece 1 prints "AA\nAA\n".
func TestPrint_SinglePiece(t *testing.T) {
	g := solver.Grid{{1, 1}, {1, 1}}
	got := captureStdout(func() { Print(g) })
	want := "AA\nAA\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestPrint_EmptyCell verifies cell value 0 renders as '.'.
func TestPrint_EmptyCell(t *testing.T) {
	g := solver.Grid{{1, 0}, {0, 2}}
	got := captureStdout(func() { Print(g) })
	want := "A.\n.B\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestPrint_LetterMapping verifies pieces map to A, B, C...
func TestPrint_LetterMapping(t *testing.T) {
	g := solver.Grid{{1}, {2}, {3}}
	got := captureStdout(func() { Print(g) })
	want := "A\nB\nC\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestPrint_AllPeriods verifies a grid of zeros prints all '.'.
func TestPrint_AllPeriods(t *testing.T) {
	g := solver.Grid{{0, 0}, {0, 0}}
	got := captureStdout(func() { Print(g) })
	want := "..\n..\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestPrint_RowNewlines verifies every row ends with '\n'.
func TestPrint_RowNewlines(t *testing.T) {
	g := solver.Grid{{1, 2}, {3, 4}}
	got := captureStdout(func() { Print(g) })
	rows := []string{}
	cur := ""
	for _, ch := range got {
		cur += string(ch)
		if ch == '\n' {
			rows = append(rows, cur)
			cur = ""
		}
	}
	if len(rows) != 2 {
		t.Errorf("expected 2 rows terminated by \\n, got %d", len(rows))
	}
	for i, row := range rows {
		if row[len(row)-1] != '\n' {
			t.Errorf("row %d does not end with \\n: %q", i, row)
		}
	}
}
