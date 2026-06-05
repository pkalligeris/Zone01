package parser

import (
	"bufio"
	"fmt"
	"os"
)

// Tetromino represents a parsed and validated tetris piece.
// Cells holds the (row, col) offsets of each of the 4 blocks,
// normalised so the top-left occupied cell is (0,0).
type Tetromino struct {
	Cells [4][2]int
}

// Parse reads the file at filePath, parses and validates all tetrominoes,
// and returns the ordered slice of Tetromino values.
// It returns a non-nil error under any invalid condition defined in the PRD.
func Parse(filePath string) ([]Tetromino, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}
	defer f.Close()

	// Read all lines from the file.
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Split lines into raw blocks separated by exactly one blank line.
	blocks, err := splitBlocks(lines)
	if err != nil {
		return nil, err
	}
	if len(blocks) == 0 {
		return nil, fmt.Errorf("file contains no tetrominoes")
	}
	if len(blocks) > 26 {
		return nil, fmt.Errorf("too many tetrominoes: max 26, got %d", len(blocks))
	}

	// Validate and parse each block.
	tetrominoes := make([]Tetromino, 0, len(blocks))
	for i, block := range blocks {
		t, err := parseBlock(block, i)
		if err != nil {
			return nil, err
		}
		tetrominoes = append(tetrominoes, t)
	}
	return tetrominoes, nil
}

// splitBlocks splits the lines of the file into groups of 4 lines,
// each group separated by exactly one blank line.
// It enforces that no two consecutive blank lines appear, and that
// there is no trailing blank line at the end of the file.
func splitBlocks(lines []string) ([][]string, error) {
	if len(lines) == 0 {
		return nil, nil
	}

	// A trailing newline from the file will manifest as a final empty string.
	// We allow exactly one trailing empty line (the natural line terminator),
	// but not two (which would be a trailing blank line between/after blocks).
	// Strip exactly one trailing empty line if present.
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	// After stripping one, a second trailing empty line is invalid.
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		return nil, fmt.Errorf("invalid spacing: trailing blank line at end of file")
	}

	var blocks [][]string
	var current []string

	for i, line := range lines {
		if line == "" {
			// A blank line is a separator — current block must be complete.
			if len(current) != 4 {
				return nil, fmt.Errorf("invalid spacing: blank line after %d rows (expected 4)", len(current))
			}
			blocks = append(blocks, current)
			current = nil

			// Check the next line is not also blank (double blank line).
			if i+1 < len(lines) && lines[i+1] == "" {
				return nil, fmt.Errorf("invalid spacing: consecutive blank lines found")
			}
		} else {
			current = append(current, line)
		}
	}

	// Handle the last block (no trailing blank line after it).
	if len(current) > 0 {
		if len(current) != 4 {
			return nil, fmt.Errorf("invalid dimensions: last block has %d rows (expected 4)", len(current))
		}
		blocks = append(blocks, current)
	}

	return blocks, nil
}

// parseBlock validates the 4 rows of a raw block for character and dimension
// correctness, then delegates shape validation to validateShape.
func parseBlock(rows []string, blockIndex int) (Tetromino, error) {
	if len(rows) != 4 {
		return Tetromino{}, fmt.Errorf("block %d: expected 4 rows, got %d", blockIndex+1, len(rows))
	}
	for r, row := range rows {
		if len(row) != 4 {
			return Tetromino{}, fmt.Errorf("block %d row %d: expected 4 cols, got %d", blockIndex+1, r+1, len(row))
		}
		for _, ch := range row {
			if ch != '.' && ch != '#' {
				return Tetromino{}, fmt.Errorf("block %d row %d: invalid character %q", blockIndex+1, r+1, ch)
			}
		}
	}

	// Validate block count, contiguity and normalise.
	return validateShape(rows, blockIndex)
}

// validateShape checks that the block has exactly 4 '#' cells that are
// contiguous, then returns the normalised Tetromino.
func validateShape(rows []string, blockIndex int) (Tetromino, error) {
	// Collect all '#' positions.
	var cells [][2]int
	for r, row := range rows {
		for c, ch := range row {
			if ch == '#' {
				cells = append(cells, [2]int{r, c})
			}
		}
	}

	if len(cells) != 4 {
		return Tetromino{}, fmt.Errorf("block %d: expected 4 '#' cells, got %d", blockIndex+1, len(cells))
	}

	// BFS contiguity check.
	if !isContiguous(rows, cells) {
		return Tetromino{}, fmt.Errorf("block %d: '#' cells are not contiguous", blockIndex+1)
	}

	return normalise(cells), nil
}

// isContiguous performs a BFS from the first '#' cell and checks
// that all 4 '#' cells are reachable.
func isContiguous(rows []string, cells [][2]int) bool {
	// Build a set of '#' positions for O(1) lookup.
	type pos struct{ r, c int }
	hashSet := make(map[pos]bool, 4)
	for _, cell := range cells {
		hashSet[pos{cell[0], cell[1]}] = true
	}

	visited := make(map[pos]bool, 4)
	queue := []pos{{cells[0][0], cells[0][1]}}
	visited[queue[0]] = true

	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			next := pos{cur.r + d[0], cur.c + d[1]}
			if hashSet[next] && !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	return len(visited) == 4
}

// normalise shifts the cell positions so the top-left occupied cell is (0,0).
func normalise(cells [][2]int) Tetromino {
	minR, minC := cells[0][0], cells[0][1]
	for _, cell := range cells {
		if cell[0] < minR {
			minR = cell[0]
		}
		if cell[1] < minC {
			minC = cell[1]
		}
	}

	var t Tetromino
	for i, cell := range cells {
		t.Cells[i] = [2]int{cell[0] - minR, cell[1] - minC}
	}
	return t
}
