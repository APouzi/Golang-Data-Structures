package randomlc

// You are given a string s and an integer array indices of the same length. The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

// Return the shuffled string.

//  s = "codeleet", indices = [4,5,6,7,0,2,1,3]
// Output: "leetcode"

func RunRandomlC(){

}

func ShuffleString(s string, indices []int)string{
	var newS []rune = make([]rune, len(indices))

	for i:=0;i<len(s);i++{
		newS[indices[i]] = rune(s[i])
	}
	var newSR string = string(newS)
	return newSR
}

//You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

// Return the max sliding window.
//Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]

func maxSlidingWindow(arr []int, k int)[]int{
	return []int{}
}