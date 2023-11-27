package patterns

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

func GetSkylinePrac(buildings [][]int)[]int{
	return buildings[0]
}
