package sorting

func QuickSort(arr []int, low, high int) {
	if low < high {

		pivot := Partition(arr, low, high)

		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

func Partition(arr []int, low, high int) int {
	greater := low - 1
	pivot := arr[high]
	for i := low; i < high; i++ {
		if arr[i] >= pivot {
			greater++
			arr[i], arr[greater] = arr[greater], arr[i]
		}
	}
	greater++
	arr[greater], arr[high] = arr[high], arr[greater]
	return greater
}
