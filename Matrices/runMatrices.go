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

	//Practice
	fmt.Println("\n---Practice problems have started---")
	
	spMatPrac := [][]int{
		{1, 2, 3,},
		{12, 13, 5},
		{10, 9, 7},
	}
	fmt.Println("Spiral Matrix Practice:",SpiralMatrixPrac(spMatPrac))

	matrixPrac2 := [][]int{
		{1, 2, 3,},
		{12, 13, 5},
		{10, 9, 7},
	}
	fmt.Println("Rotate Matrix Practice:", RotateMatrixPrac(matrixPrac2))

	// matrix3Prac := [][]int{
	// 	{ 1,  2, 3,},
	// 	{12, 0,  5},
	// 	{10, 9,  7},
	// }
	matrix3Prac2 := [][]int{
		{ 0,  2, 0,},
		{12, 9,  5},
		{0, 9,  0},
	}
	fmt.Println("Set Matrix Zero Practice", SetMatrixZerosPrac(matrix3Prac2))

	matrixPrac4 := [][]int{
		{ 1, 2, 3, 3,},
		{12, 0, 4, 5},
		{10, 9, 7, 7},
	}
	fmt.Println("mirror matrix:", MirrorMatrixPrac(matrixPrac4))

	matrixBallPrac := [][]int{
		{1,1,1,-1,-1},
		{1,1,1,-1,-1},
		{-1,-1,-1,1,1},
		{1,1,1,1,-1},
		{-1,-1,-1,-1,-1}}

	fmt.Println("where will the ball drop?", WhereWillTheBallDropPrac(matrixBallPrac))
}
