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
		startIndex = i //startIndex will constanstly be set to the current i for us to insert into the stack, unless we need to set it to the new one that will be replaced from elimination.
		for len(monostack) > 0 && v < monostack[len(monostack)-1][0] {
			var popped [2]int = monostack[len(monostack)-1]
			maxArea = max(maxArea, (i-popped[1])*popped[0])//This is the important portion of the entire algo. What we are saying here, is that while the stack is holding a value  bigger than the current height, we want pop it and then perform a calculation on it.
			//The calculation: We get the current index and subtract it from the popped pair and get that index. That index represents the "startIndex" as in where that height "represents" all the coverage is can get. Because remember, that area can stretch all the way to start.
			//Lastly, we multiply it by the stacks poppede height. so example, "2" index is 4, and we pop he top of stack "6", with index of 3. so (4-3) * 6. Then we pop again and get "5" with index of "2", (4-2) * 5. That will be the max area.
			startIndex = popped[1]//Lastly, we need to assign the startIndex to the popped index. This is because while popping while current height is smaller than top of stack, we know that it can cover the entirety of what's being popped because it's smaller than the popped heights. So we replace startIndex "i" with the popped one.
			monostack = monostack[:len(monostack)-1]
		}
		monostack = append(monostack, [2]int{v, startIndex})
	}
	//Lastly, we need to make sure that the rest of the stack isn't holding out on better maxArea.
	for _, pair := range monostack {
		//Since we already went through the entire list, we get the entire length of heights and substract it from the popped index.
		maxArea = max(maxArea, pair[0]*(len(heights)-pair[1]))
	}
	return maxArea
}