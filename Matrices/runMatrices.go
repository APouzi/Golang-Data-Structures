package matrices

import "fmt"

func RunMatrix() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}
	fmt.Println("Spiral Matrix result",SpiralMatrix(matrix))
	
	matrix2 := [][]int{
		{1, 2, 3,},
		{12, 13, 5},
		{10, 9, 7},
	}
	fmt.Println("Rotate Matrix:", RotateMatrix(matrix2))

	matrix3 := [][]int{
		{ 1,  2, 3,},
		{12, 0,  5},
		{10, 9,  7},
	}

	fmt.Println("Set Zeros",setZeroes(matrix3))

	matrix4 := [][]int{
		{ 1, 2, 3, 3,},
		{12, 0, 4, 5},
		{10, 9, 7, 7},
	}
	fmt.Println("mirror matrix:", MirrorMatrixLeftToRight(matrix4))
	
	matrixBall := [][]int{
		{1,1,1,-1,-1},
		{1,1,1,-1,-1},
		{-1,-1,-1,1,1},
		{1,1,1,1,-1},
		{-1,-1,-1,-1,-1}}

	fmt.Println("where will the ball drop?",whereTheBallDrops(matrixBall))
}
