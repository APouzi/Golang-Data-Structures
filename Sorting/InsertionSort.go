package sorting

func InsertionSortAlgo(arr []int) {
	var sorted int
	var key int
	for i := 1; i < len(arr); i++ {
		sorted = i - 1
		key = arr[i]
		for sorted >= 0 && key < arr[sorted] {
			arr[sorted+1] = arr[sorted]
			//This is shifting everything to the right. The first to change will always be the key that will be shifted into. This is going to be
			sorted--
			//We are ALWAYS going down from the key, which is moving down.
		}
		arr[sorted+1] = key // We must always do a sorted + 1 because at first we reached -1 and need to be the first element of the array. Doing the above will always move the sorted portion of the array one before that we want to do it.

	}
}