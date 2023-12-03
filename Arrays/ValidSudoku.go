package arrays

//Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules

//Input: board = 
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: true
func ValidSudoku(board [][]byte) bool {
	var row, col map[int]map[byte]bool = make(map[int]map[byte]bool), make(map[int]map[byte]bool)
	var box map[[2]int]map[byte]bool = make(map[[2]int]map[byte]bool)

	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]bool)
		col[i] = make(map[byte]bool)
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			if board[r][c] == 46 {
				continue
			}
			var rc [2]int = [2]int{r / 3, c / 3}

			if row[r][board[r][c]] || col[c][board[r][c]] || box[rc][board[r][c]] {
				return false
			}
			if box[rc] == nil {
				box[rc] = make(map[byte]bool)
			}
			row[r][board[r][c]], col[c][board[r][c]], box[rc][board[r][c]] = true, true, true
		}
	}
	return true
}