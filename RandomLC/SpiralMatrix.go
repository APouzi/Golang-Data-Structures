package randomlc

import "fmt"

// {1,   2, 3},
// {10, 11, 4},
// {9,  12, 5},
// {8,   7, 6},
func SpiralMatrix(arr [][]int){
	var rowStart,colStart,colEnd, rowEnd int = 0,0,len(arr[0])-1, len(arr)-1
	
	fmt.Println("\nRow and Col:",rowStart,colStart,rowEnd,colEnd)
	ans := []int{}
	for rowStart < rowEnd && colStart < rowEnd{
		for i:=colStart; i <= colEnd; i++ {
			ans = append(ans, arr[rowStart][i])
		}
		rowStart++
		for i := rowStart; i <=rowEnd; i++{
			ans = append(ans,arr[i][colEnd])
		}
		colEnd--
		for i :=colEnd; i >= colStart; i--{
			ans = append(ans, arr[rowEnd][i])
		}
		rowEnd--
		for i := rowEnd; i >= rowStart; i--{
			ans = append(ans, arr[i][colStart])
		}
		colStart++
	}

	fmt.Println(ans)

}