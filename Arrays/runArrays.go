package arrays

import "fmt"

func RunArrays() {
	fmt.Println("---Array Problems Have Begun---")
	fmt.Println("Shuffle String:", ShuffleString("codeleet", []int{4,5,6,7,0,2,1,3}))
	fmt.Println("Product of Array Except itself:", ProductExceptSelf([]int{1,2,3,4}))
	fmt.Println("Validate Sudoku:", ValidSudoku([][]byte{{'5','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},{'.','9','8','.','.','.','.','6','.'},{'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},{'.','.','.','.','8','.','.','7','9'}}))
	fmt.Println("Group Anagrams", GroupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
	fmt.Println("FindKthLargest", FindKthLargest([]int{3,2,1,5,6,4},2))
	fmt.Println("Longest Consecutive Sequence", LongestConsecutive([]int{100,4,200,1,3,2}))
	fmt.Println("Missing Number",MissingNumber([]int{3,0,1}))
	fmt.Println("Product Except Self", []int{1,2,3,4})
	fmt.Println("Top K Frequent",TopKFrequent([]int{1,1,1,2,2,3},2))
	fmt.Println("Two Sum",TwoSum([]int{2,7,11,15},9))
	fmt.Println("Is Anagram",IsAnagramStringTrick("hello","llohe"))


	//Practice Problems
	fmt.Println("----Practice Array Problems---")
	fmt.Println("Shuffle String Practice:", ShuffleString("codeleet", []int{4,5,6,7,0,2,1,3}))
	fmt.Println("Group Anagrams Practice",GroupAnagramsPract([]string{"eat","tea","tan","ate","nat","bat"}))
	fmt.Println("FindKthLargest Practice",FindKthLargestPrac([]int{3,2,1,5,6,4},2))
	fmt.Println("Longest Consecutive Sequence Practice",LongestConsercutiveSequencePrac([]int{0,3,7,2,5,8,4,6,0,1}))
	fmt.Println("Missing Number Practice", MissingNumberPrac([]int{3,0,1}))
	fmt.Println("Product Except Self Pracitce",ProductExceptSelfPrac([]int{1,2,3,4}))
	fmt.Println("Top K Frequent Practice", TopKFrequentPrac([]int{1,1,1,2,2,3},2))
	fmt.Println("Two Sum Practice")
	fmt.Println("Is Anagram Practice",IsAnamgramPrac("hello","olelh"))
	fmt.Println("Valid Sudoku Practice", ValidSudokuPrac([][]byte{{'5','3','.','.','7','.','.','.','.'},{'6','.','.','1','9','5','.','.','.'},{'.','9','8','.','.','.','.','6','.'},{'8','.','.','.','6','.','.','.','3'},{'4','.','.','8','.','3','.','.','1'},{'7','.','.','.','2','.','.','.','6'},{'.','6','.','.','.','.','2','8','.'},{'.','.','.','4','1','9','.','.','5'},{'.','.','.','.','8','.','.','7','9'}}))
	fmt.Println("Shuffle String",ShuffleStringPrac("codeleet", []int{4,5,6,7,0,2,1,3}))
}