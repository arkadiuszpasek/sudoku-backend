package sudoku

import (
  "math/rand/v2"
)

func LoopIncrement(value int) int {
	if value == 9 {
		return 1
	}
	return value + 1
}

func GenerateBoard() BoardValues {
	base := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(base), func(i int, j int) {
		base[i], base[j] = base[j], base[i]
	})
	
	boardValues := make(BoardValues, 9)
	boardValues[0] = base

	for i := 1; i < 9; i++ {
		boardValues[i] = make([]int, 9)

		for j := 0; j < 9; j++ {
			boardValues[i][j] = LoopIncrement(boardValues[i-1][j])
		}
	}

	rand.Shuffle(len(boardValues), func(i int, j int) {
		boardValues[i], boardValues[j] = boardValues[j], boardValues[i]
	})

	return boardValues
}
