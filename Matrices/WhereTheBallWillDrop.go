package matrices

func whereTheBallDrops(arr [][]int) []int {
	colEnd, rowEnd := len(arr[0])-1, len(arr)-1
	ans := []int{}
	for i := 0; i <= colEnd; i++ {
		curr_col := i
		for row := 0; row <= rowEnd; row++ {
			//Here we are asking if the column is going to be going left or right. Say our curr_col is 0, we then want to add curr_col to the arr[row][curr_col] which will be +1. That will now mean that the next row will have us on column 1.
			next_col := curr_col + arr[row][curr_col]

			//before we move down the row to column 1, we need to test if:
			//1) next_col is bigger than 0 aka not on the left side of the wall
			//2) next_col is smaller than the column size aka not past the right side of the wall
			//3) there isn't a v shape in the column that will trap the ball
			//If any of these are true, means the ball is trapped and we need to assign the curr_col as -1 and break out the nested loop to go to next iteration
			if next_col < 0 || next_col >= colEnd || arr[row][curr_col] != arr[row][next_col] {
				curr_col = -1
				break
			}
			//assigned the current column we are on as the next column because that will be where the next row will end up.
			curr_col = next_col
		}
		//once we reach the last row, it means the ball has dropped out in it's location and we need to assign that the array or tray of balls
		ans = append(ans, curr_col)
	}
	return ans
}