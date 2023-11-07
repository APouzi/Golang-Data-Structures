package sorting

func QuickSort(arr []int, low, high int) {
	if low < high {

		pivot := Partition(arr, low, high) //The pivot, like mid in merge is going to be the divider, except that the pivot will be in the correct place when done so we need to use this as the divider and get the pivot (+-1) and recuservly do this until we are done.

		//The actual array isn't being divided and then merged together into a new array. This is just getting the source array and the indexs will be altering or "dividing" until we reach elements that are the size of one. Once they are that size, they are technically completed
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

func Partition(arr []int, low, high int) int {
	greater := low
	pivot := arr[high]
	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			arr[i], arr[greater] = arr[greater], arr[i]
			greater++
		}
	}
	arr[greater], arr[high] = arr[high], arr[greater]
	return greater
}

// In textbooks it's like this for some reason, 
func LumotoPartition(arr []int, low, high int) int {
	greater := low - 1
	pivot := arr[high]
	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			greater++
			arr[i], arr[greater] = arr[greater], arr[i]
		}
	}
	arr[greater+1], arr[high] = arr[high], arr[greater+1]
	return greater + 1
}
