package matrices

func whereTheBallDrops(arr [][]int) []int {
	colEnd, rowEnd := len(arr[0])-1, len(arr)-1
	ans := []int{}
	for i := 0; i <= colEnd; i++ {
		curr_col := i
		for j := 0; j <= rowEnd; j++ {
			next_col := curr_col + arr[j][curr_col]

			if next_col < 0 || next_col >= colEnd || arr[j][curr_col] != arr[j][next_col] {
				curr_col = -1
				break
			}
			curr_col = next_col
		}
		ans = append(ans, curr_col)
	}
	return ans
}