package main

import (
	"sudoku"
)

func generateAndPrintBoard() {
	board := sudoku.GenerateBoard()
	sudoku.ConsolePrintBoard(board)
}

func main() {
	generateAndPrintBoard()
}
