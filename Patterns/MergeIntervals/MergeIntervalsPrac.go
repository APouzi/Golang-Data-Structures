package mergeintervals

import (
	"container/heap"
	"sort"
)

//Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

// Example 1:
// Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
// Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

// Example 2:
// Input: intervals = [[1,4],[4,5]]
// Output: [[1,5]]
// Explanation: Intervals [1,4] and [4,5] are considered overlapping.

//Note, DO NOT use quicksort for intervals, this is because quicksort intervals is unstable.
func MergeIntervalsPrac(intervals [][]int) [][]int {
	return [][]int{}
}

//You are given an array of non-overlapping intervals intervals where intervals[i] = [starti, endi] represent the start and the end of the ith interval and intervals is sorted in ascending order by starti. You are also given an interval newInterval = [start, end] that represents the start and end of another interval.
//Insert newInterval into intervals such that intervals is still sorted in ascending order by starti and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary).
//Return intervals after the insertion.

//Example 1:
//Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
//Output: [[1,5],[6,9]]

//Example 2:
//Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//Output: [[1,2],[3,10],[12,16]]
//Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
func InsertIntervalsPrac(intervals [][]int, newInterval []int) [][]int {
	return [][]int{}
}


//Given an array of intervals intervals where intervals[i] = [starti, endi], return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
// Example 1:
// Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
// Output: 1
// Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.

// Example 2:
// Input: intervals = [[1,2],[1,2],[1,2]]
// Output: 2
// Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.

// Example 3:
// Input: intervals = [[1,2],[2,3]]
// Output: 0
// Explanation: You don't need to remove any of the intervals since they're already non-overlapping.
//NOTE as example 3 shows, if two edges are the same, they are NOT overlapping.
// Input: intervals = [[1,2],[2,3],[3,4],[1,3]] Output: 1 [[1,2],[1,3],[2,3],[3,4]]
// Input: intervals = [[1,2],[1,2],[1,2]] Output: 2
func nonOverLappingIntervalsPrac(intervals [][]int) int {
	return -1
}


//A city's skyline is the outer contour of the silhouette formed by all the buildings in that city when viewed from a distance. Given the locations and heights of all the buildings, return the skyline formed by these buildings collectively.

// The geometric information of each building is given in the array buildings where:
//buildings[i] = [lefti, righti, heighti]:
// lefti is the x coordinate of the left edge of the ith building.
// righti is the x coordinate of the right edge of the ith building.
// heighti is the height of the ith building.

// You may assume all buildings are perfect rectangles grounded on an absolutely flat surface at height 0.
// The skyline should be represented as a list of "key points" sorted by their x-coordinate in the form [[x1,y1],[x2,y2],...]. Each key point is the left endpoint of some horizontal segment in the skyline except the last point in the list, which always has a y-coordinate 0 and is used to mark the skyline's termination where the rightmost building ends. Any ground between the leftmost and rightmost buildings should be part of the skyline's contour.

// Note: There must be no consecutive horizontal lines of equal height in the output skyline. For instance, [...,[2 3],[4 5],[7 5],[11 5],[12 7],...] is not acceptable; the three lines of height 5 should be merged into one in the final output as such: [...,[2 3],[4 5],[12 7],...]

//Example 1
// buildings = [[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
// Output: [[2,10],[3,15],[7,12],[12,0],[15,10],[20,8],[24,0]]

// Example 2:
// buildings = [[0,2,3],[2,5,3]]
// Output: [[0,3],[5,0]]
type MaxHeap2 [][]int

func (h MaxHeap2) Len() int           { return len(h) }
func (h MaxHeap2) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h MaxHeap2) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap2) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MaxHeap2) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func GetSkylinePrac(buildings [][]int)[][]int{
	var ans [][]int = [][]int{}
	var points []int = make([]int, len(buildings)*2)
	for i,v := range buildings{
		points[i*2] = v[0]
		points[i*2+1] = v[1]
	}
	sort.Slice(points, func(i, j int) bool {return points[i] < points[j]})
	var unique []int = []int{points[0]}
	for i:=1;i<len(points);i++{
		if points[i-1] != points[i]{
			unique = append(unique, points[i])
		}
	}
	var building, prevHeight, currHeight int = 0, 0 ,0
	h := &MaxHeap2{}
	heap.Init(h)
	for _,point := range unique{
		
		for h.Len() > 0 && point >= (*h)[0][1]{
			heap.Pop(h)
		}

		for building < len(buildings) && point >= buildings[building][0]{
			heap.Push(h, []int{buildings[building][2],buildings[building][1]})
			building++
		}

		currHeight = 0
		if h.Len() >0{
			currHeight = (*h)[0][0]
		}

		if currHeight != prevHeight{
			ans = append(ans, []int{point,currHeight})
			prevHeight = currHeight
		}
	}
	return ans
}

// You are given a 2D integer array intervals, where intervals[i] = [lefti, righti] describes the ith interval starting at lefti and ending at righti (inclusive). The size of an interval is defined as the number of integers it contains, or more formally righti - lefti + 1. 
//You are also given an integer array queries. The answer to the jth query is the size of the smallest interval i such that lefti <= queries[j] <= righti. If no such interval exists, the answer is -1.

// Return an array containing the answers to the queries.

// Example 1:
// Input: intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
// Output: [3,3,1,4]

func MinIntervalPrac(intervals [][]int, queries []int)[]int{
	return []int{}
}

//Given an array of meeting time intervals consisting of start and end times[[s1,e1],[s2,e2],...](si< ei), determine if a person could attend all meetings.
//Input:[[0,30],[5,10],[15,20]]
//Output: false
func MeetingRoomsPrac(meetings [][]int)bool{
	return false
}

//Given an array of meeting time intervals consisting of start and end times[[s1,e1],[s2,e2],...](si< ei), find the minimum number of conference rooms required.
func MeetingRoomsIIPrac(intervals [][]int) int{
	return -1
}