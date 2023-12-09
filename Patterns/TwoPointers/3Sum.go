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
	sort.Slice(nums, func(i, j int) bool { //First thing is we want to actually sort the sums, this would allow us to easily tell if there is a repeat by by checking if the current num is the same as the last one. This also allows us to use the pointer technique to either move pointer to the left if result too big, to the right if result is too small. [-4,-1,-1,0,1,2]
		return nums[i] < nums[j]
	})
	var mid, right int //Since technically, we have 3 pointers. The left, mid and right pointer.
	var ans [][]int = [][]int{}
	for left := 0; left < len(nums)-2; left++ { //Since we want to start from the left, we need to stop with the accounting of mid being next to left and right being the end, like: [0,0,1,left,mid,right]
		if left > 0 && nums[left] == nums[left-1] {//When we sorted the nums, we gave the ability to see if there was a copy of the last.
			continue
		}  
		mid = left + 1
		right = len(nums) - 1
		for mid < right {
			if nums[left]+nums[mid]+nums[right] > 0 { //if the left + mid + right is bigger than 0, decrease the right pointer.
				right--
			} else if nums[left]+nums[mid]+nums[right] < 0 {//if the left + mid + right is smaller than 0, increase the mid point.
				mid++
			} else {
				//If we are equal to 0, we need to append this answer.
				ans = append(ans, []int{nums[left], nums[mid], nums[right]})
				//The problem is that if we don't move the mid pointer up, (apparently it doesn't matter if we move the mid or the right pointer). Then we get stuck in the loop because nothing changes.
				mid++
				for nums[mid] == nums[mid-1] && mid < right { //As long as mid and mid-1 are the same, then we move mid up while making sure it will never pass right.
					mid++
				}
			}

		}
	}
	return ans
}