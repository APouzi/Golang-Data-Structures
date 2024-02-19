package arrays

import (
	"container/heap"
)

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Example 1:
// Input: nums = [3,2,3]
// Output: 3

// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2


func MajorityElement(nums []int) int {
    var majority int = nums[0]
    var count int = 1

    for _, v := range nums{
        if majority != v{
            count--
        }else{
            count++
        }
        if count == 0{
            majority = v
            count = 1
        }
    }

    return majority
}





type majority [][]int 

func(m majority) Len() int{return len(m)}
func(m majority) Less(i,k int) bool{return m[i][0] > m[k][0]}
func(m majority) Swap(i,k int) {m[i],m[k] = m[k], m[i]}
func(m *majority) Pop() any{
    old := *m
    n := len(old)
    popped:= old[n-1]
    *m = old[0:n-1]
    return popped
}
func (m *majority) Push(x any){
    *m = append(*m, x.([]int))
}

func MajorityElementHeap(nums []int) int {
    h := &majority{}
    heap.Init(h)
    var maj map[int]int = map[int]int{}

    for _, v := range nums{
        maj[v]++
        heap.Push(h, []int{maj[v],v})
    }

    return (*h)[0][1]
}