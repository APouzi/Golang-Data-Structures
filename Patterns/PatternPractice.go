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

func MaxSlidingWindowPrac(nums []int, k int) []int{
	return []int{}
}

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

// Notice that the solution set must not contain duplicate triplets.

// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.


func ThreeSum(nums []int) [][]int {
	return [][]int{}
}

//You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]). Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

//Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

// Example 2:
// Input: height = [1,1]
// Output: 1

func ContainerWithMostWaterPrac(height []int) int {
	var left, right int = 0, len(height)-1
	var maxArea int = 0
	for left < right {
		if height[left] < height[right]{
			maxArea = max(maxArea, min(height[left],height[right]) * (right - left))
			left++
		}else{
			maxArea = max(maxArea, min(height[left],height[right]) * (right - left))
			right--
		}
	}
	return maxArea
}

//Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

// Example 1:
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.

// Example 2:
// Input: height = [4,2,0,3,2,5]
// Output: 9


func TrappingRainWater(height []int) int {
	return -1
}

//Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 < numbers.length.

// Return the indices of the two numbers, index1 and index2, ADDED BY ONE as an integer array [index1, index2] of length 2.

// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.

// Example 1:
// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2] (NOT [0,1])
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

// Example 2:
// Input: numbers = [2,3,4], target = 6
// Output: [1,3]
// Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

// Example 3:
// Input: numbers = [-1,0], target = -1
// Output: [1,2]
// Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
func TwoSumIIPrac(numbers []int, target int) []int {
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
	return -1
}

//You are given an m x n integer matrix matrix with the following two properties:

// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if target is in matrix or false otherwise.

// You must write a solution in O(log(m * n)) time complexity.

// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true

//Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//Output: false


func SearchMatrix(matrix [][]int, target int) bool {
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
	return -1
}