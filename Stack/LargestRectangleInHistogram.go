package stack

// Given an array of integers heights representing the histogram's bar height where the width of each bar is 1, return the area of the largest rectangle in the histogram.

// Example 1:
// Input: heights = [2,1,5,6,2,3]
// Output: 10
// Explanation: The above is a histogram where width of each bar is 1.
// The largest rectangle is shown in the red area, which has an area = 10 units.


func largestRectangleArea(heights []int) int {
	var monostack [][2]int = make([][2]int, 0)
	var maxArea int = 0
	var startIndex int
	for i, v := range heights {
		startIndex = i
		for len(monostack) > 0 && v < monostack[len(monostack)-1][0] {
			var popped [2]int = monostack[len(monostack)-1]
			maxArea = max(maxArea, (i-popped[1])*popped[0])
			startIndex = popped[1]
			monostack = monostack[:len(monostack)-1]
		}
		monostack = append(monostack, [2]int{v, startIndex})
	}
	for _, pair := range monostack {
		maxArea = max(maxArea, pair[0]*(len(heights)-pair[1]))
	}
	return maxArea
}