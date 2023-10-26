package sorting

import "fmt"

func MergeSort(array []int, st int, end int) {
	if st < end {
		mid := (st + end) / 2 
		MergeSort(array, st, mid) //First break up the left side of the array, recursively and this will be put into the stack until we reach the base case of "st < end" being the size of 0, 1 is the smallest
		MergeSort(array, mid+1, end) //Pop it out the stack and then break up the rest, to the right. 
		// Note for the two top 
		Merge(array, st, mid, end) //We then merge the left and right arrays. This process will repeat over and over because of the fact that the stack at the bottom won't until the very last steps. The bottom will have the two biggest size of the arrays.
	}
}

func Merge(array []int, st int, mid int, end int) {
	fmt.Println("inputs st, mid, end", st, mid, end)
	aLL := mid - st + 1 //This is just getting the size of left array 
	aLR := end - mid // This is just getting the size of left array  
	fmt.Println(aLL,aLR)
	arrL := make([]int, aLL) //Here we don't use []int{}, unless we want to do appending operation, which would be more costly
	arrR := make([]int, aLR) //Here we don't use []int{}, unless we want to do appending operation, which would be more costly
	
}