package binarysearch

//Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:
// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4

// Example 2:
// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1

func BinarySearch(nums []int, target int) int {
	var left, right int = 0, len(nums) - 1
	var middle int
	for left <= right {
		middle = (left + right) / 2
		if nums[middle] < target { //[-1,0,3,5[Middle],9,12]
			//Here we find out that the middle is smaller than the target, this means we have to move the left to the middle plus 1.
			left = middle + 1
		} else if nums[middle] > target {//Here we see that the middle is bigger than the target, so we can remove all the ones from the right. This means that we can completely ignore all the right of the array starting at the middle as canidates.
			right = middle - 1
		} else if nums[middle] == target {
			return middle
		}
	}
	return -1
}