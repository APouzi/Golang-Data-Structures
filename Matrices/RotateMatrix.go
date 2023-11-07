package matrices

//You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

//You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

//Traverse groups of four cells in the matrix, starting from the four corners
//swap the top left cell with the top right cell
//swap the top left cell with the bottom right cell
//swap the top left cell with the bottom left cell
//Move to the next group of four cells and repeat the above steps

func RotateMatrix(arr [][]int) [][]int{

	quadrants := len(arr)/2
	for i:=0; i < quadrants; i++{
		var RowStart, ColStart, RowEnd, ColEnd int = 0+i,0+i,len(arr)-1-i,len(arr[0])-1-i
		var rotations int = len(arr)-1-(i*2)
		for j:=0; j<rotations;j++{
			topLeft := arr[RowStart][ColStart+j]
			arr[RowStart][ColStart+j] = arr[RowEnd-j][ColStart]
			arr[RowEnd-j][ColStart] = arr[RowEnd][ColEnd-j]
			arr[RowEnd][ColEnd-j] = arr[RowStart+j][ColEnd]
			arr[RowStart+j][ColEnd] = topLeft
		}
	}
	return arr
}