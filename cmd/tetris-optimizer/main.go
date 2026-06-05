package main

import (
	"fmt"
	"os"

	"tetris-optimizer/internal/parser"
	"tetris-optimizer/internal/printer"
	"tetris-optimizer/internal/solver"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	pieces, err := parser.Parse(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	grid, err := solver.Solve(pieces)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(1)
	}

	printer.Print(grid)
}
