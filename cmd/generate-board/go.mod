module generate-board

go 1.24.0

require (
	github.com/aws/aws-lambda-go v1.54.0
	sudoku v0.0.0
)

replace sudoku => ../../sudoku
