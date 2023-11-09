package sorting

//Best case O(n)
//average O(n2)
//worst case O(n2)
func InsertionSortAlgo(arr []int) {
	var sorted int
	var key int
	for i := 1; i < len(arr); i++ {
		sorted = i - 1
		key = arr[i]
		// I ran into a problem before where I had just used a singular "if statement" (for key < arr[sorted]) instead of putting it inside of the for loop, this wont work because sometimes we have to cut this shift over short. example:
		//7,6,3,1,5,4, key = 5, i = 1. What happens with this is that key will have to stop when the 3 shifts to the right and 5 needs to take it's place and not go further.
		for sorted >= 0 && key < arr[sorted] {
			arr[sorted+1] = arr[sorted]
			//This is shifting everything to the right. The first to change will always be the key that will be shifted into. This is going to be
			sorted--
			//We are ALWAYS going down from the key, which is moving down.
		}
		arr[sorted+1] = key // We must always do a sorted + 1 because at first we reached -1 and need to be the first element of the array. Doing the above will always move the sorted portion of the array one before that we want to do it.

	}
}