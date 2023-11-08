package matrices

// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's. You must do it in place.
//Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
//Output: [[1,0,1],[0,0,0],[1,0,1]]
func setZeroes(matrix [][]int) [][]int {
	rowEnd, colEnd := len(matrix)-1, len(matrix[0])-1
	var col map[int]bool = map[int]bool{}
	var row map[int]bool = map[int]bool{}
	for y := 0; y <= rowEnd; y++ {
		for x := 0; x <= colEnd; x++ {
			if matrix[y][x] == 0 {
				col[y] = true
				row[x] = true
			}
		}
	}

	for y := 0; y <= rowEnd; y++ {
		for x := 0; x <= colEnd; x++ {
			if col[y] || row[x] {
				matrix[y][x] = 0
			}
		}
	}
	return matrix
}