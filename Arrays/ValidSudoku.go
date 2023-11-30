package arrays

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