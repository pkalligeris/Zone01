package printer

import (
	"fmt"

	"tetris-optimizer/internal/solver"
)

// Print renders the solved grid to os.Stdout.
// Each cell value 0 is printed as '.'; values 1..26 map to 'A'..'Z'.
// Each row is terminated by a newline character.
func Print(g solver.Grid) {
	for _, row := range g {
		for _, val := range row {
			if val == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(string(rune('A' + val - 1)))
			}
		}
		fmt.Print("\n")
	}
}
