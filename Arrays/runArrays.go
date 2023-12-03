package arrays

import "fmt"

func RunArrays() {
	fmt.Println("---Array Problems Have Begun---")
	fmt.Println("Shuffle String:", ShuffleString("codeleet", []int{4,5,6,7,0,2,1,3}))
	fmt.Println("Product of Array Except itself:", ProductExceptSelf([]int{1,2,3,4}))
	fmt.Println("Validate Sudoku:", ValidSudoku([][]byte{{'5','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},{'.','9','8','.','.','.','.','6','.'},{'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},{'.','.','.','.','8','.','.','7','9'}}))
	fmt.Println("Group Anagrams")
	fmt.Println("FindKthLargest")
	fmt.Println("Longest Consecutive Sequence")
	fmt.Println("Missing Number")
	fmt.Println("Product Except Self Pracitce")
	fmt.Println("Top K Frequent")
	fmt.Println("Two Sum")
	fmt.Println("Is Anagram")


	//Practice Problems
	fmt.Println("----Practice Array Problems---")
	fmt.Println("Shuffle String Practice:", ShuffleString("codeleet", []int{4,5,6,7,0,2,1,3}))
	fmt.Println("Group Anagrams Practice")
	fmt.Println("FindKthLargest Practice")
	fmt.Println("Longest Consecutive Sequence Practice")
	fmt.Println("Missing Number Practice")
	fmt.Println("Product Except Self Pracitce")
	fmt.Println("Top K Frequent Practice")
	fmt.Println("Two Sum Practice")
	fmt.Println("Is Anagram Practice")
	fmt.Println("Valid Sudoku Practice")
}