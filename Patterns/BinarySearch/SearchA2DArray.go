package binarysearch

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
	var rows, columns int = len(matrix) - 1, len(matrix[0]) - 1
	var top, bottom int = 0, rows
	var middleRow int
	for top <= bottom {
		middleRow = (bottom + top) / 2
		if target < matrix[middleRow][0] {
			bottom = middleRow - 1
		} else if target > matrix[middleRow][columns] {
			top = middleRow + 1
		} else {
			break
		}
	}
	if top <= bottom == false {
		return false
	}

	var left, right int = 0, columns
	var middle int
	for left <= right {
		middle = (right + left) / 2
		if target > matrix[middleRow][middle] {
			left = middle + 1
		} else if target < matrix[middleRow][middle] {
			right = middle - 1
		} else if matrix[middleRow][middle] == target {
			return true
		}
	}
	return false
}

func SearchMatrixV2(matrix [][]int, target int) bool {
	var rows, columns int = len(matrix) - 1, len(matrix[0]) - 1
	var top, bottom int = 0, rows
	var middleRow int
	for top <= bottom {
		middleRow = (bottom + top) / 2
		if target < matrix[middleRow][0] {
			bottom = middleRow - 1
		} else if target > matrix[middleRow][columns] {
			top = middleRow + 1
		} else {
			return BinarySearchRow(matrix[middleRow], target)
		}
	}
	return false
}

func BinarySearchRow(row []int, target int) bool {
	var left, right int = 0, len(row) - 1
	var middle int
	for left <= right {
		middle = (right + left)/2
		if target < row[middle] {
			right = middle - 1
		} else if target > row[middle] {
			left = middle + 1
		} else if row[middle] == target {
			return true
		}
	}
	return false
}