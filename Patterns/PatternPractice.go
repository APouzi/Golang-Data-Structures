package patterns

import (
	"container/heap"
	"sort"
)

// Write an algorithm to determine if a number n is happy.
// A happy number is a number defined by the following process:
// Starting with any positive integer, replace the number by the sum of the squares of its digits.Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.

//Example 1:
// Input: n = 19
// Output: true

// Example 2:
// Input: n = 2
// Output: false
func HappyNumberPrac(num int) bool {
	return false
}

func cycleHN(num int) int {
	return -1
}

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


// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Example 1:
// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

// Example 2:
// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.
func BuyHighSellLowPrac(prices []int) int{
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





//Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.

// Example 1:
// Input: nums = [-1,0,3,5,9,12], target = 9
// Output: 4
// Explanation: 9 exists in nums and its index is 4

// Example 2:
// Input: nums = [-1,0,3,5,9,12], target = 2
// Output: -1
// Explanation: 2 does not exist in nums so return -1

func BinarySearch(nums []int, target int) int {
	var left, right int = 0 , len(nums)-1
	var mid int
	for left <= right{
		mid = (left + right)/2
		if target > nums[mid]{
			left = mid+1
		}else if target < nums[mid]{
			right = mid -1
		}else if target == nums[mid]{
			return mid
		}
	}
	return -1
}

//You are given an m x n integer matrix with the following two properties:

// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.

// You must write a solution in O(log(m * n)) time complexity.

// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true

//Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//Output: false
//[1, 3, 5, 7],
//[10,11,16,20],
//[23,30,34,60]
//target = 3

func SearchMatrix(matrix [][]int, target int) bool {
	var bottom, top int = 0 ,len(matrix)-1
	var arr int
	var left, right int = 0, len(matrix[0])-1
	for bottom <= top{
		arr = (bottom + top)/2
		if target < matrix[arr][left] {
			top = arr - 1
		}else if target > matrix[arr][right] {
			bottom = arr + 1
		}else{
			break
		}
	}
	
	var mid int
	for left <= right{
		mid = (right+left)/2
		if matrix[arr][mid] < target{
			left = mid + 1
		}else if matrix[arr][mid] > target{
			right = mid - 1
		}else{
			return true
		}
	}
	return false
}

// Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

// Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

// Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

// Return the minimum integer k such that she can eat all the bananas within h hours.

// Example 1:
// Input: piles = [3,6,7,11], h = 8
// Output: 4

// Example 2:
// Input: piles = [30,11,23,4,20], h = 5
// Output: 30

// Example 3:
// Input: piles = [30,11,23,4,20], h = 6
// Output: 23


func MinEatingSpeed(piles []int, h int) int {
	var left, right int = 1, maxPile(piles)
	var mid float64
	var h64 float64 = float64(h)
	var hours float64
	var ans int = right
	for left <= right{
		mid = float64((left + right)/2)
		for _, p := range piles{
			hours += float64(p)/mid
		}
		if mid <= h64{
			ans = min(ans, int(mid))
			right = int(mid) -1 
		}else{
			left = int(mid) + 1
		}
	}
	return ans
}

func maxPile(piles []int)int{
	var ans int = piles[0]
	for _, v := range piles{
		if v > ans{
			ans = v
		}		
	}
	return ans
}



//You are given two lists of closed intervals, firstList and secondList, where firstList[i] = [starti, endi] and secondList[j] = [startj, endj]. Each list of intervals is pairwise disjoint and in sorted order.

// Return the intersection of these two interval lists.

// A closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.

// The intersection of two closed intervals is a set of real numbers that are either empty or represented as a closed interval. For example, the intersection of [1, 3] and [2, 4] is [2, 3].

// Example 1:
// Input: firstList = [[0,2],[5,10],[13,23],[24,25]], secondList = [[1,5],[8,12],[15,24],[25,26]]
// Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]

// Example 2:
// Input: firstList = [[1,3],[5,9]], secondList = []
// Output: []


func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0{
		return [][]int{}
	}

	var lists [][]int = [][]int{}

	lists = append(lists, firstList...)
	lists = append(lists, secondList...)
	var endPoints [][]int = make([][]int, len(lists) *2)
	for i,v := range lists{
		endPoints[2 * i] = []int{v[0],1}
		endPoints[2 * i + 1] = []int{v[1],-1}
	}

	sort.Slice(endPoints, func(i, j int) bool {
		if endPoints[i][0] == endPoints[j][0]{
			return endPoints[i][1] > endPoints[j][1]
		}
		return endPoints[i][0] < endPoints[j][0]})
	var ans [][]int = [][]int{}
	var delta, startPoint int = 0, 0
	for _, v := range endPoints{
		delta += v[1]

		if delta == 2{
			startPoint = v[0]
		}

		if delta == 1 && v[1] == -1  {
			ans = append(ans, []int{startPoint, v[0]})
		}
		
	}
	return ans
}

// Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

// [4,5,6,7,0,1,2] if it was rotated 4 times.
// [0,1,2,4,5,6,7] if it was rotated 7 times.
// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

// Given the sorted rotated array nums of unique elements, return the minimum element of this array.

// You must write an algorithm that runs in O(log n) time.

// Example 1:
// Input: nums = [3,4,5,1,2]
// Output: 1
// Explanation: The original array was [1,2,3,4,5] rotated 3 times.

// Example 2:
// Input: nums = [4,5,6,7,0,1,2]
// Output: 0
// Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

// Example 3:
// Input: nums = [11,13,15,17]
// Output: 11
// Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

//25,26,27,30,1,3,4,6,7,8,9,11,13,14,27
func findMinInRotatedArrayPrac(nums []int) int {
	return -1
}


