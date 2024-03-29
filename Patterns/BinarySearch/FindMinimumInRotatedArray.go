package binarysearch

// Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

// [4,5,6,7,0,1,2] if it was rotated 4 times.
// [0,1,2,4,5,6,7] if it was rotated 7 times.
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

// Given the sorted rotated array nums of unique elements, return the minimum element of this array.

// You must write an algorithm that runs in O(log n) time.

// Example 1:
// Input: nums = [3,4,5,1,2]
// Output: 1
// Explanation: The original array was [1,2,3,4,5] rotated 3 times.

// Example 2:
// Input: nums = [4,5,6,7,0,1,2]
// Output: 0
// Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

// Example 3:
// Input: nums = [11,13,15,17]
// Output: 11
// Explanation: The original array was [11,13,15,17] and it was rotated 4 times.


func FindMinInRotatedArray(nums []int) int {
	var left, right int = 0, len(nums) - 1
	var mid int
	//4,5,6,7,0,1,2
	for left < right {
		mid = (left + right)/2
        if nums[mid] > nums[right]{
            left = mid + 1
        }else if nums[mid] < nums[left]{
            right = mid
        }else{
            break
        }
		//In other binary searches, we are using the middle to find the target, in this specific problem, what we are doing is we are using the edges/guardrails/left&right to get the answer. Because we are looking for a maxima/minima.
	}
	return nums[left] //both right and left can be used. By time we are done, both left and right will be left over!
}