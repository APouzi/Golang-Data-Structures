package sorting

import "fmt"

func RunSorting() {
	arr := []int{5, 4, 3, 2, 1}
	InsertionSort := []int{9,8,7,6,5,4,3,2,1,11,12,34,64,7654,123,1236,124,1235,643,67,123,565,3123}
	// arr2 := []int{5, 4, 3, 2, 1}
	BubbleSort(arr)
	// BubbleSort2(arr2)
	fmt.Println("BubbleSort",arr)
	InsertionSortAlgo(InsertionSort)
	fmt.Println("Insertion Sort",InsertionSort)
}