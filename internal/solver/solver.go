package solver

import (
	"math"

	"tetris-optimizer/internal/parser"
)

// Grid is the mutable square board used during backtracking.
// 0 means empty; values 1..N map to tetrominoes A..Z.
type Grid [][]int

// NewGrid allocates and returns a size×size Grid filled with 0.
func NewGrid(size int) Grid {
	g := make(Grid, size)
	for i := range g {
		g[i] = make([]int, size)
	}
	return g
}

// MinSize returns the minimum square side length needed to hold n tetrominoes.
// Formula: ⌈√(n × 4)⌉
func MinSize(n int) int {
	return int(math.Ceil(math.Sqrt(float64(n * 4))))
}

// Solve finds the smallest square grid that fits all pieces using backtracking.
// It starts at MinSize and grows the grid by 1 until a solution is found.
func Solve(pieces []parser.Tetromino) (Grid, error) {
	size := MinSize(len(pieces))
	for {
		g := NewGrid(size)
		if place(pieces, 0, g, size) {
			return g, nil
		}
		size++
	}
}

// place attempts to place piece at index into the grid recursively.
// Returns true if all pieces are successfully placed.
func place(pieces []parser.Tetromino, index int, g Grid, size int) bool {
	if index == len(pieces) {
		return true
	}

	piece := pieces[index]
	label := index + 1

	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if canPlace(piece, r, c, g, size) {
				// Place the piece.
				for _, cell := range piece.Cells {
					g[r+cell[0]][c+cell[1]] = label
				}

				if place(pieces, index+1, g, size) {
					return true
				}

				// Undo placement.
				for _, cell := range piece.Cells {
					g[r+cell[0]][c+cell[1]] = 0
				}
			}
		}
	}

	return false
}

// canPlace checks whether a piece can be placed at (startR, startC) on the grid.
func canPlace(piece parser.Tetromino, startR, startC int, g Grid, size int) bool {
	for _, cell := range piece.Cells {
		r := startR + cell[0]
		c := startC + cell[1]
		if r < 0 || r >= size || c < 0 || c >= size {
			return false
		}
		if g[r][c] != 0 {
			return false
		}
	}
	return true
}
