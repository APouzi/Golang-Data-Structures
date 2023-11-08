package matrices

func MirrorMatrixLeftToRight(arr [][]int) [][]int {
	var colStart, rowStart, colEnd int = 0, 0, len(arr[0])-1
	var colLim, rowLim int = len(arr[0])/2 , len(arr)
	for i := 0; i < colLim; i++ {
		for j := 0; j < rowLim; j++ {
			arr[rowStart+j][colStart+i], arr[rowStart+j][colEnd-i] = arr[rowStart+j][colEnd-i], arr[rowStart+j][colStart+i]
		}
	}
	return arr
}