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

func FindDuplicateInArrayPrac([]int)int{
	return -1
}

func GetSkylinePrac(buildings [][]int)[][]int{
	return buildings
}


// Given a string(str1), is str2 a permutation of str1?
// Example str1 = "abcdefg", str2 = "bac". "bac" is a permuatation of "abcdefg" because "abc" of "abcdefg" can be permutated to "bac"
// test cases:
// str1 = "eidbaooo", str2 = "ab"
// str1 = "eidboaoo", str2 = "ab"
// str1 = "eidboaoo", str2 "eidboaoo"

func PermutationOfStringPrac(str1, str2 string) bool{
	return false
}


// Given an integer array nums, find a
// subarray
//  that has the largest product, and return the product.
// The test cases are generated so that the answer will fit in a 32-bit integer.

// Example 1:
// Input: nums = [2,3,-2,4]
// Output: 6
// Explanation: [2,3] has the largest product 6.

// Example 2:
// Input: nums = [-2,0,-1]
// Output: 0
// Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
func MaxProductPrac(nums []int) int {
	return -1
}

//Given an array of integers nums and an integer k, return the number of contiguous subarrays where the product of all the elements in the subarray is strictly less than k.

// Example 1:
// Input: nums = [10,5,2,6], k = 100
// Output: 8
// Explanation: The 8 subarrays that have product less than 100 are:
// [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6]
// Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.

// Example 2:
// Input: nums = [1,2,3], k = 0
// Output: 0

func NumSubarrayProductLessThanKPrac(nums []int, k int) int {
	return -1
}

//Given a string s, find the length of the longest substring without repeating characters.

// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

//testcases: "qrsvbspk" = 5, " " = 1
//first one will fail without a for loop inside deleting instead of if loop.

func LengthOfLongestSubstringPrac(s string) int {
	return -1
}

// You are visiting a farm that has a single row of fruit trees arranged from left to right. The trees are represented by an integer array fruits where fruits[i] is the type of fruit the ith tree produces.You want to collect as much fruit as possible. However, the owner has some strict rules that you must follow:

// You only have two baskets, and each basket can only hold a single type of fruit. There is no limit on the amount of fruit each basket can hold.
// Starting from any tree of your choice, you must pick exactly one fruit from every tree (including the start tree) while moving to the right. The picked fruits must fit in one of your baskets.
// Once you reach a tree with fruit that cannot fit in your baskets, you must stop.
// Given the integer array fruits, return the maximum number of fruits you can pick.

// Example 1:
// Input: fruits = [1,2,1]
// Output: 3
// Explanation: We can pick from all 3 trees.

// Example 2:
// Input: fruits = [0,1,2,2]
// Output: 3
// Explanation: We can pick from trees [1,2,2].
// If we had started at the first tree, we would only pick from trees [0,1].

// Example 3:
// Input: fruits = [1,2,3,2,2]
// Output: 4
// Explanation: We can pick from trees [2,3,2,2].
// If we had started at the first tree, we would only pick from trees [1,2].

func TotalFruitPrac(fruits []int) int {
	return -1
}

// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character. You can perform this operation at most k times.
// Return the length of the longest substring containing the same letter you can get after performing the above operations.

// Example 1:
// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.

// Example 2:
// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.
// There may exists other ways to achieve this answer too.

func CharacterReplacementPrac(s string, k int) int {
	return -1
}


// Given two strings s and t of lengths m and n respectively, return the minimum window
// substring of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
// The testcases will be generated such that the answer is unique.

// Example 1:
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

// Example 2:
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.

// Example 3:
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.

//testcases s = "acbbaca" , t = "aba"
func MinWindowPrac(s string, t string) string {
	return "practice"
}