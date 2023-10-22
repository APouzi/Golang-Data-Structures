package sorting

func InsertionSortAlgo(arr []int) {
	var sorted int
	var key int
	for i := 1; i < len(arr); i++ {
		sorted = i - 1
		key = arr[i]
		for sorted >= 0 && key < arr[sorted] {
			arr[sorted+1] = arr[sorted]
			sorted--
		}
		arr[sorted+1] = key

	}
}