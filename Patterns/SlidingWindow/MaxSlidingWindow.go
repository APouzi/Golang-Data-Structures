package slidingwindow

import (
	"container/heap"
)

//My attempt but way too slow:
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