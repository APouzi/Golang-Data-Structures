package arrays

import "container/heap"

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]

// Example 2:
// Input: nums = [1], k = 1
// Output: [1]

//nums =[5,3,1,1,1,3,73,1]
// k = 2

type MinHeap [][]int 

func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Swap(i, j int) {h[i],h[j]=h[j],h[i]}
func (h MinHeap) Less( i, j int) bool { return h[i][0] < h[j][0]}

func (h *MinHeap) Push(i interface{}){
    *h = append(*h, i.([]int))
}

func (h *MinHeap) Pop()interface{}{
    old := *h
    oldL := len(old)
    last := old[:oldL-1]
    *h = old[0:oldL-1]
    return last
}

func TopKFrequent(nums []int, k int) []int {
    h := &MinHeap{}
    heap.Init(h)
    count := map[int]int{}
    for i := 0; i < len(nums);i++{
        count[nums[i]]++
    }
    for i, v := range count{
        heap.Push(h, []int{v,i})
    }
    ans := []int{}
    for k > 0{
        k--
        ans = append(ans,heap.Pop(h).([]int)[1])
    }
    
    return ans
}