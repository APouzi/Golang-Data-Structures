package sorting

func MergeSort(array []int, st int, end int) {
	if st < end {
		mid := (st + end) / 2
		MergeSort(array, st, mid)    //First break up the left side of the array, recursively and this will be put into the stack until we reach the base case of "st < end" being the size of 0, 1 is the smallest
		MergeSort(array, mid+1, end) //Pop it out the stack and then break up the rest, to the right.
		// Note for the two top
		Merge(array, st, mid, end) //We then merge the left and right arrays. This process will repeat over and over because of the fact that the stack at the bottom won't until the very last steps. The bottom will have the two biggest size of the arrays.
	}
}

func Merge(array []int, st int, mid int, end int) {
	// fmt.Println("inputs st, mid, end", st, mid, end)
	aLL := mid - st + 1 //This is just getting the size of left array
	aLR := end - mid    // This is just getting the size of left array
	// fmt.Println(aLL,aLR)

	arrL := make([]int, aLL) //Here we don't use []int{}, unless we want to do appending operation, which would be more costly
	arrR := make([]int, aLR) //Here we don't use []int{}, unless we want to do appending operation, which would be more costly

	// Lets build out the new arrays to compare to and then reorder in our source array
	for i := 0; i < aLL; i++ {
		arrL[i] = array[st+i]
	}

	for i := 0; i < aLR; i++ {
		arrR[i] = array[i+mid+1]
	}
	// fmt.Println("Left", arrL,"Right", arrR)

	// initiate the left and right array's tracking of indexs as we compare the two arrays. Also intiate the start of the sourceArray's insertion index, which starts on the recursive's start.
	i, j, arri := 0, 0, st

	// Compare each array to one another until both or one of the array's have no more to compare to.
	for i < aLL && j < aLR {
		if arrL[i] <= arrR[j] {
			array[arri] = arrL[i]
			i++
			arri++
		} else {
			array[arri] = arrR[j]
			j++
			arri++
		}
	}

	// If one of the arrays have something left, empty out the rest, left going first and then right.
	for i < aLL {
		array[arri] = arrL[i]
		i++
		arri++
	}

	for j < aLR {
		array[arri] = arrR[j]
		j++
		arri++
	}
}


func MergeSortLinkedList(head *Node) *Node{
	if head == nil || head.next == nil{
		return head
	}
	
	mid := MiddleLL(head)
	left := MergeSortLinkedList(head)
	right := MergeSortLinkedList(mid)
	return MergeLL(left, right)
}


}