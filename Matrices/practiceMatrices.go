package matrices

// {1,  2,  3,  4},
// {12, 13, 14, 5},
// {11, 16, 15, 6},
// {10, 9,  8,  7},
func SpiralMatrixPrac(arr [][]int) []int {

	return []int{}
}

//rotate image 90 degrees
//  >
// {1, 2, 3,},\/
// {12, 13, 5},
//^{10, 9, 7},
//          <

func RotateMatrixPrac(arr [][]int)[][]int{
	quadrants := len(arr)/2
	var rowStart, colStart, rowEnd, colEnd int = 0, 0, len(arr)-1, len(arr[0])-1
	for q := 0; q < quadrants; q++{
		steps := len(arr) - 1 - (q*2)
		for i := 0; i < steps; i++{
			topLeft := arr[rowStart][colStart+i]
			arr[rowStart][colStart+i] = arr[rowEnd-i][colStart]
			arr[rowEnd-i][colStart] = arr[rowEnd][colEnd-i]
			arr[rowEnd][colEnd-i] = arr[rowStart+i][colEnd]
			arr[rowStart+i][colEnd] = topLeft
		}
	}
	return arr
}

//if there is a zero, turn all of the values in it's row and column into 0s
func SetMatrixZerosPrac(arr [][]int)[][]int{
	

	return arr

}

//Mirror the matrix array
func MirrorMatrixPrac(arr [][]int)[][]int{
	return [][]int{}
}
//You have a 2-D grid of size m x n representing a box, and you have n balls. The box is open on the top and bottom sides.

// Each cell in the box has a diagonal board spanning two corners of the cell that can redirect a ball to the right or to the left.

// A board that redirects the ball to the right spans the top-left corner to the bottom-right corner and is represented in the grid as 1.
// A board that redirects the ball to the left spans the top-right corner to the bottom-left corner and is represented in the grid as -1.
// We drop one ball at the top of each column of the box. Each ball can get stuck in the box or fall out of the bottom. A ball gets stuck if it hits a "V" shaped pattern between two boards or if a board redirects the ball into either wall of the box.

// Return an array answer of size n where answer[i] is the column that the ball falls out of at the bottom after dropping the ball from the ith column at the top, or -1 if the ball gets stuck in the box.
//NOTE: the last row is not the floor but the last portion of the slide

//Input: grid = [
// [1,1,1,-1,-1],
// [1,1,1,-1,-1],
// [-1,-1,-1,1,1],
// [1,1,1,1,-1],
// [-1,-1,-1,-1,-1]]
// Output: [1,-1,-1,-1,-1]

// Input: grid = [
// [1,  1, 1, 1, 1, 1],
// [-1,-1,-1,-1,-1,-1],
// [ 1, 1, 1, 1, 1, 1],
// [-1,-1,-1,-1,-1,-1]]
// Output: [0,1,2,3,4,-1]

//Given values of a that represent a left leaning wall (-1) and right leaning wall (1), where will the ball drop, even if it can?
func WhereWillTheBallDropPrac(arr [][]int)[]int{
	var ans []int = []int{}
	for col := 0; col < len(arr[0]);col++{
		cur_col := col
		for row := 0; row < len(arr);row++{
			next_col := cur_col + arr[row][cur_col]

			if next_col < 0 || next_col > len(arr[0])-1 || arr[row][next_col] != arr[row][cur_col]{
				cur_col = -1
				break
			}

			cur_col = next_col
		}
		ans = append(ans, cur_col)
	}
	return ans
}


