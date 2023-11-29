package patterns

import "sort"

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

func PermuStrPrac(str1 string, str2 string) bool {
	return false
}

func MergeIntervalsPrac(intervals [][]int) [][]int {
	return [][]int{}
}

// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
func InsertIntervalsPrac(intervals [][]int, newInterval []int) [][]int {
	return [][]int{}
}

// Input: intervals = [[1,2],[2,3],[3,4],[1,3]] Output: 1 [[1,2],[1,3],[2,3],[3,4]]
// Input: intervals = [[1,2],[1,2],[1,2]] Output: 2
func nonOverLappingIntervalsPrac(intervals [][]int) int {
	return -1
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

func BuyHighSellLowPrac(prices []int) int{
	return -1
}

func ValidPalindromePrac(word string) bool{
	return false
}

func FindDuplicatePrac([]int)int{
	return -1
}

func GetSkylinePrac(buildings [][]int)[][]int{
	var pointX []int = make([]int, len(buildings)*2)
	var heap HeapPairMax = HeapPairMax{}

	for i := 0; i < len(buildings);i++{
		pointX[i*2] = buildings[i][0]
		pointX[i*2+1] = buildings[i][1]
	}
	
	var uniqueX []int = []int{pointX[0]}
	for i :=1;i<len(pointX);i++{
		if pointX[i-1] != pointX[i]{
			uniqueX = append(uniqueX, pointX[i])
		}
	}
	sort.Ints(uniqueX)
	var b int = 0
	var prevMaxHeight int = 0
	var skyline [][]int
	for x := 0; x <len(uniqueX);x++{

		for b <len(buildings) && buildings[b][0] <= uniqueX[x]{
			heap.Insert([]int{buildings[b][2],buildings[b][1]})
			b++
		}

		for len(heap.arr) > 0 && heap.arr[0][1] <= uniqueX[x]{
			heap.Pop()
		}

		var currentHeight int = 0
		if len(heap.arr) >0{
			currentHeight = heap.arr[0][0]
		}

		if currentHeight != prevMaxHeight{
			skyline = append(skyline, []int{uniqueX[x],currentHeight})
			prevMaxHeight = currentHeight
		}
	}


	return skyline
}
