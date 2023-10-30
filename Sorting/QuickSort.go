package sorting

func QuickSort(arr []int, low, high int) {
	if low < high {

		pivot := Partition(arr, low, high)

		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

