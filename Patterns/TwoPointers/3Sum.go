package twopointers

import "sort"

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.


func ThreeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var mid, right int
	var ans [][]int = [][]int{}
	for left := 0; left < len(nums)-2; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		mid = left + 1
		right = len(nums) - 1
		for mid < right {

			if nums[left]+nums[mid]+nums[right] > 0 {
				right--
			} else if nums[left]+nums[mid]+nums[right] < 0 {
				mid++
			} else {
				ans = append(ans, []int{nums[left], nums[mid], nums[right]})
				mid++
				for nums[mid] == nums[mid-1] && mid < right {
					mid++
				}
			}

		}
	}
	return ans
}