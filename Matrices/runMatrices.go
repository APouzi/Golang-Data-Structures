package matrices

import "fmt"

func RunMatrix() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	SpiralMatrix(matrix)
	
	fmt.Println("Spiral Matrix Practice:",SpiralMatrixPrac(matrix))
	matrix2 := [][]int{
		{1, 2, 3,},
		{12, 13, 5},
		{10, 9, 7},
	}
	fmt.Println("Rotate Matrix:", RotateMatrix(matrix2))
}
