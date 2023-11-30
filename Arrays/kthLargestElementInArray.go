package arrays

func findKthLargest(nums []int, k int) int {
	h := Heap{}
	for _, v := range nums {
		h.Insert(v)
		if len(h.arr) > k {
			h.Pop()
		}
	}

	return h.arr[0]
}