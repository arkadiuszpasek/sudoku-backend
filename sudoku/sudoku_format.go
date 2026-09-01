package sudoku

import "fmt"

func ConsolePrintBoard(boardValues BoardValues) {
	for i, row := range boardValues {
		if i % 3 == 0 && i != 0 {
			fmt.Println("------+-------+------")
		}
		for j, cell := range row {
			if j % 3 == 0 && j != 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", cell)
		}
		fmt.Println()
	}
}