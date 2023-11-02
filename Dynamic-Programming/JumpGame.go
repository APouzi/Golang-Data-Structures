package dynamicprogramming

func canJump(nums []int) bool {
	var LastGoodIndex int = len(nums) - 1
	var size int = len(nums) - 1
	for i := size; i >= 0; i-- {
		if i+nums[i] >= LastGoodIndex {
			LastGoodIndex = i
		}
	}
	return LastGoodIndex == 0
}