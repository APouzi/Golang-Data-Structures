package monotonicdeque

//Monotonic Deques or Arrays are in two possible states:
//Either the array is growing or the array is decreasing (as in the x[i] > x[i+1] ) or the array is increasing(x[i] < x[i+1]) or the stack is staying the same.

//Given the list of [2,1,2,4,3], find the NEXT (to the right) greater hieght. so 0th element (2) next greater height is going to be 4, 1th element(1) is going to be 2, 3nd element(2) is going to be 4, 3rd element(4) is going to be nothing or -1 because it has no greater element. Since we want the next greater stack, we need to reverse iterate over this array and add each element to the monotonic stack in reverse order
func MonotonicArrayInit(nums []int) []int {
	var mono []int = []int{}
	var ans []int = make([]int, len(nums))

	for i := len(nums)-1; i >= 0; i-- {
		for len(mono) > 0 && nums[i] >= mono[len(mono)-1] {
			mono = mono[:len(mono)-1]
		}

		if len(mono) == 0 {
			ans[i] = -1
		}else{
			ans[i] = mono[len(mono)-1]
		}
		mono = append(mono, nums[i])
	}
	return ans
}