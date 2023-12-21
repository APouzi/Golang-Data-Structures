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
		//[0,1,0,2,1,0,1,3,2,1,2,1]
		//>>>>>>>> 		 <<<<<<<<<<
		//We are getting max left and right, and when we check the trapped water, we are only checking the one that is ahead of the max left and right. NOT the one behind it.
		if leftMax < rightMax {//This is asking which boundary is the smallest out of the two. We ONLY want the smallest of the two as that is what matters in retaining water.
			left++
			leftMax = max(height[left], leftMax) // As we go forward, we need to be asking which "wall" is the heighest. This is because with this answer, we are able to ascertain the ith biggest to the left. With that knowledge, we are able to get the biggest wall. This answer will only change when we find the next biggest wall. What is import is that we are only doing calculations based on the walls that are lower or equal to it to get the water that each isolated unit can hold.

			ans += leftMax - height[left] //These will will never be negative, the reason why is because leftMax is only able to reach on the same value as height[left] when it updates and never less. leftMax can only be bigger or the same value as height[left] which would amount to 0. leftMax can never be "ahead" of left, just at the same level. Same applies to rightMax.
		} else { 
			right--
			rightMax = max(height[right], rightMax)
			ans += rightMax - height[right]
		}
	}
	return ans
}

//-------The Why requires understanding another algorithm-------

//The above algorithm is the replacement of the algorithmm below:

//LeftMax: One loop starts from the left or 0th index and saves the state of the maxima at each iteration
//RightMax: One loop starts from the right or last index and goes to the 0th index and saves the state of the maxima at each iteration
//MinLR: The third loop will get the min of the LeftMax and RightMax and store at each iteration.
//Lastly, we loop over the original list of heights and subtract the MinLR from heights[i]. Then increment that result.
//LeftMax:	  [0,1,1,2,2,2,2,3,3,3,3,3]
//RightMax:	  [1,2,2,2,3,3,3,3,3,3,3,3]
//MinLR:   	  [0,1,1,2,2,2,2,3,3,3,3,3]
//Last Loop:  ans += heights[i] - MinLR[i] 

//With our version, we are replacing the need for four loops and instead, only moving up when the leftMax/rightMax is smaller than it's counterpart. That way we achieve the same exact thing as leftMax and rightMax loops, but only doing when one is smaller than the other. 