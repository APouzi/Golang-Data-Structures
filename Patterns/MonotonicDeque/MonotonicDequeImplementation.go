package monotonicdeque

func MonotonicArrayInit(nums []int) []int {
	var mono []int = []int{}

	for i := 0; i < len(nums); i++ {
		if len(mono) > 0 && nums[i] > mono[len(mono)-1] {
			mono = mono[:len(mono)-1]
		}

		if len(mono) == 0 {
			mono = append(mono, nums[i])
		}
	}
	return mono
}