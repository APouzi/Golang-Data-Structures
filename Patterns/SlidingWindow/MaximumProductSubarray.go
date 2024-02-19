package slidingwindow

// Given an integer array nums, find a
// subarray
//  that has the largest product, and return the product.
// The test cases are generated so that the answer will fit in a 32-bit integer.

// Example 1:
// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.

// Example 2:
// Input: nums = [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

//Other test cases: [7,-2,-4] = 56
func MaxProduct(nums []int) int {
	var ans int = maxOne(nums) //This does two things: It get's the max singular value in the list in case thats the max we can get. Two, if 0 is the largest, then it will be the max, because we skip them in the "v==0" check

	//We are setting currMin and currMax to 1, because we need to multiply the first number by 1.
	var currMin, currMax, temp int = 1, 1,1
	for _, v := range nums{
		if v == 0{
            currMin, currMax = 1, 1 //If v is 0, then it will forever be 0 and we can't break out if it. So we skip this and we set these to 1's. So it's like we start the subarray at the next value. Basically saying the max and min are the next value starting at that position.
            continue
        }
        temp = currMax * v//Set the temp to currMax * v, because otherwise currMin will introduce a bug because our currMax will be changed entirely.
        currMax = max(v * currMax, v * currMin,  v)//This will be the leader that will find your answer. 
		currMin = min(v * currMin, temp, v)//This is a stop measure to make sure that what you multiplied together isn't smaller than the current value, the last max, or the currMin * v. What this does, is it helps you with the negative numbers.
		//7,-2,-4, by time we get to -4, the currMax is 7 and currMin is -14. Then we have the choice of figuring out if -14(currMin) * -4 is bigger than 7(currMax) * -4 or is just negative 4.
        ans = max(ans, currMax)//Here we keep track of the biggest and when we get past -4, the choice is either 7 or 56 (-14 * 4)
	}
 	return ans  
}

func maxOne(nums []int) int{
    var ret int = nums[0]
    for _, v := range nums{
        ret = max(v, ret)
    }
    return ret
}   