package matrices

//You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

//You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.



func RotateMatrix(arr [][]int) [][]int{

//Traverse groups of four cells in the matrix, starting from the four corners
//swap the top left cell with the top right cell
//swap the top left cell with the bottom right cell
//swap the top left cell with the bottom left cell
//Move to the next group of four cells and repeat the above steps

	quadrants := len(arr)/2
	for i:=0; i < quadrants; i++{
		var RowStart, ColStart, RowEnd, ColEnd int = 0+i,0+i,len(arr)-1-i,len(arr[0])-1-i
		var rotations int = len(arr)-1-(i*2)//This is an important one to remember because what we are saying is that for every quadrant, we have a set amount of rotations. The rotations are judged as the size of array and we substract by (i*2) and this makes sense if you look at it visually. 
		//if the array is the size is 5x5
		//   1  2  3  4  5
		//   6  7  8  9 10
		//  11 12 13 14 15
		//  16 17 18 19 20
		//  21 22 23 24 25
		// 1) Quadrants: "len(arr)/2 aka 4/2" = 2
		// 2) roatations: len(arr) -1 - (i*2):
		//		*First set: 1, 2, 3, 4, 5  (5 - 1 - (0*2)) = 4
		//		*Second set: 	7, 8, 9    (5 - 1 - (1*2)) = 3
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