package slidingwindow

import (
	"container/heap"
)

//You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position.

// Return the max sliding window.

//Example 1:
// Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
// Output: [3,3,5,5,6,7]
// Explanation:
// Window position                Max
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
//  1 [3  -1  -3] 5  3  6  7       3
//  1  3 [-1  -3  5] 3  6  7       5
//  1  3  -1 [-3  5  3] 6  7       5
//  1  3  -1  -3 [5  3  6] 7       6
//  1  3  -1  -3  5 [3  6  7]      7

// Example 2:
// Input: nums = [1], k = 1
// Output: [1]



func MaxSlidingWindow(nums []int, k int)[]int{

	//Here we are going to be creating a monotonic array to be allowing to keep track of the biggest element by using a decreasing monotonic stack/array. Decreasing = [8,4,3,1,].
	//We need to do this because what we want to be able to get 0th index, which should be the absolute largest index in a subarray. 
	var mono, ans []int = make([]int, 0), make([]int, 0)//ans can also be make([]int, len(nums) - k+1) to get window size and make it run more effiecent.

	var l int = 0
	for r := 0; r < len(nums);r++{

		//If the current iteration of nums[r] is bigger than the top of the stack, we want to pop all the elements that are smaller than the current iteration of nums[r]. Once we pop all elements that are smaller than the elements in the stack, we want to insert that into the monotonic array.
		for len(mono) > 0 && nums[mono[len(mono)-1]] < nums[r]{
			mono = mono[:len(mono)-1]
		}
		//Make sure you understand that we are inputting the index. This is because of the fact that we want to keep track of the index due to keeping the window in size.
		mono = append(mono, r)
		//If left pointer is bigger than the index at the bottom of the monotonic array, which is the index of the biggest element in the subarray at the time, then we want to remove that element. This is because the window no longer applies.
		if mono[0] < l{
			mono = mono[1:]
		} 
		//If the right pointer is bigger than bigger than the k element, it means we can start appened the biggest element that is representing the subarray.
		if (r+1) >= k{
			ans = append(ans, nums[mono[0]])
			l++
		}

	}
	return ans
}





//My attempt but way too slow, especially on large sets of data:
type MaxHeap []int

func(h MaxHeap) Len()int{ return len(h)}
func(h MaxHeap) Less(i, j int)bool{ return h[i] > h[j]}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func(h *MaxHeap) Push(i any){ *h = append(*h, i.(int))}
func(h *MaxHeap) Pop()interface{}{
    old := *h
    oldLen := len(old)
    lastItem := old[oldLen-1]
    *h = old[0:oldLen-1]
    return lastItem
} 

func MaxSlidingWindowMyAttempt(nums []int, k int) []int {
    h := &MaxHeap{}
    heap.Init(h)
    listAns := []int{}
    windows := len(nums) - k +1
    
    for r := 0; r < windows; r++{
        inputs := r
        count :=0
        for inputs < len(nums) && count < k{
            heap.Push(h, nums[inputs])
            inputs++
            count++
        } 
       listAns = append(listAns, (*h)[0])
       for h.Len() >0{
           heap.Pop(h)
       }
    }
    return listAns
}