package sorting

import "fmt"

func SelectionSortAlgo(arr []int) {
	min := 0
	steps := 0

	for i := 0; i < len(arr); i++ {
		min = i //You must always move this back to the left most part of the shrinking window.
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] { // This cannot be "arr[j]<arr[i]", as sometimes I want to do because of the fact that this needs to get the smallest, comparing to "i" won't help because it will not be (always) the smallest, or biggest
				min = j
			}
			steps++
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	fmt.Println("steps for algo",steps)
}
