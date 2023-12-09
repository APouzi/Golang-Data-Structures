package twopointers

//Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

// Example 1:
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

// Example 2:
// Input: height = [4,2,0,3,2,5]
// Output: 9


func TrappingRainWater(height []int) int {

	var ans int = 0
	var left, right int = 0, len(height) - 1
	var leftMax, rightMax int = height[left], height[right]

	for left < right {
		if leftMax < rightMax {
			left++
			leftMax = max(height[left], leftMax)
			ans += leftMax - height[left]
		} else {
			right--
			rightMax = max(height[right], rightMax)
			ans += rightMax - height[right]
		}
	}
	return ans
}