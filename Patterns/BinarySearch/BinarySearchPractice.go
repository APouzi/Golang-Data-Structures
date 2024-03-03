package binarysearch

import "sort"

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

func BinarySearchPrac(nums []int, target int) int {
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

func SearchMatrixPrac(matrix [][]int, target int) bool {
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


func MinEatingSpeedPrac(piles []int, h int) int {
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


func IntervalIntersectionPrac(firstList [][]int, secondList [][]int) [][]int {
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