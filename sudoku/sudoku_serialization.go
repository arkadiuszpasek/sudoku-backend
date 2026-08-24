package sudoku

import (
	"fmt"
	"strings"
)

const boardSize = 9
const serializedLength = boardSize * boardSize

type BoardValues [][]int

type SudokuBoard struct {
	Board    BoardValues
	Solution BoardValues
}

func Serialize(board BoardValues) string {
	var builder strings.Builder
	builder.Grow(serializedLength)

	for _, row := range board {
		for _, cell := range row {
			builder.WriteByte(byte('0' + cell))
		}
	}

	return builder.String()
}

func Deserialize(serialized string) (BoardValues, error) {
	if len(serialized) != serializedLength {
		return nil, fmt.Errorf("serialized board must be %d digits, got %d", serializedLength, len(serialized))
	}

	board := make(BoardValues, boardSize)
	for row := 0; row < boardSize; row++ {
		board[row] = make([]int, boardSize)
		for col := 0; col < boardSize; col++ {
			char := serialized[row*boardSize+col]
			if char < '0' || char > '9' {
				return nil, fmt.Errorf("invalid digit %q at position %d", char, row*boardSize+col)
			}
			board[row][col] = int(char - '0')
		}
	}

	return board, nil
}
