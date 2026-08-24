package main

import (
	"fmt"

	"sudoku"
)

func main() {
	fmt.Println(sudoku.Serialize(sudoku.GenerateBoard()))
}
