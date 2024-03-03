package binarysearch

import "fmt"

func RunBinarySearch() {
	//---Binary Search---
	var binarySearchInput []int = []int{-1, 0, 3, 5, 9, 12}
	fmt.Println("\nBinary Search:", BinarySearch(binarySearchInput, 9))
	fmt.Println("Binary Search Practice:", BinarySearchPrac(binarySearchInput, 9))

	//---Search A 2D Array---
	var TwoDArray [][]int = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	//var TwoDArrayFalse [][]int = [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}} // target = 13
	fmt.Println("\nSearch a 2D Array", SearchMatrixV2(TwoDArray, 16))
	fmt.Println("Search a 2D Array Prac", SearchMatrixPrac(TwoDArray, 16))

	//---Koko Eating a Banana---
	var pilesBanana []int = []int{30, 11, 23, 4, 20}
	var hour int = 5
	fmt.Println("\nKoko Eating a Banana:", MinEatingSpeed(pilesBanana, hour))
	fmt.Println("Koko Eating a Banana Prac:", MinEatingSpeedPrac(pilesBanana, hour))

	//---Find Minimum In Rotated Array---
	var rotateArr []int = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("\nFind Minimum In Rotated Array:", FindMinInRotatedArray(rotateArr))
	fmt.Println("Find Minimum In Rotated Array Prac:", findMinInRotatedArrayPrac(rotateArr))
}