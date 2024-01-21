package binarysearch

func search(nums []int, target int) int {
	var left, right int = 0, len(nums) - 1
	var middle int
	for left <= right {
		middle = (left + right) / 2
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] == target {
			return middle
		}
	}
	return -1
}