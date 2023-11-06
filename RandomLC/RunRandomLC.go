package randomlc

// You are given a string s and an integer array indices of the same length. The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

// Return the shuffled string.

//  s = "codeleet", indices = [4,5,6,7,0,2,1,3]
// Output: "leetcode"

func RunRandomlC(){
	matrix := [][]int{
		{1,2,3,4},
		{12,13,14,5},
		{11,16,15,6},
		{10,9,8,7},
	}
	SpiralMatrix(matrix)
}

func ShuffleString(s string, indices []int)string{
	var newS []rune = make([]rune, len(indices))

	for i:=0;i<len(s);i++{
		newS[indices[i]] = rune(s[i])
	}
	var newSR string = string(newS)
	return newSR
}