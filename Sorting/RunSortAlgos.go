package sorting

import "fmt"

type Node struct{
	val int
	next *Node
}



func RunSorting() {
	BubbleSortNums := []int{9,8,7,6,5,4,3,2,1,11,12,34,64,7654,123,1236,124,1235,643,67,123,565,3123}
	
	// ----BubbleSort----
	BubbleSort(BubbleSortNums)
	fmt.Println("\nBubbleSort",BubbleSortNums)
	BubbleSortNums2 := []int{9,8,7,6,5,4,11,12,34,123,565,3123,123,1236,64,7654,124,1235,643,3,2,1,67}
	BubbleSort2(BubbleSortNums2)
	fmt.Println("Practice BubbleSort",BubbleSortNums2)

	//--- Insertion Sort ---
	InsertionSortNums := []int{9,8,7,6,5,4,3,2,1,11,12,34,64,7654,123,1236,124,1235,643,67,123,565,3123}
	InsertionSortAlgo(InsertionSortNums)
	fmt.Println("\nInsertion Sort",InsertionSortNums)


	// ----Selection Sort----
	SelectionSortNums := []int{9,8,7,6,5,4,3,2,1,11,12,34,64,7654,123,1236,124,1235,643,67,123,565,3123}
	SelectionSortAlgo(SelectionSortNums)
	fmt.Println("SelectionSort",SelectionSortNums)
	

	// ----------Merge Sort (Array and Linked List)----------
	fmt.Println("\n---Merge Sort Algorithm has begun---")
	MergeSortNum := []int{6,5,12,10,9,1}
	MergeSort(MergeSortNum, 0, len(MergeSortNum)-1)
	fmt.Println("MergeSort Algorithm Result",MergeSortNum)
	fmt.Println("---Merge Sort Algorithm 2 has begun---")
	MergeSortNum2 := []int{5,3,4,1,2}
	arr1 := MergeSort2(MergeSortNum2)
	fmt.Println("MergeSort2 Algorithm Result", arr1)

	fmt.Println("\n--- Merge Sort Algorithm LinkedList has begun ---")
	head, _ := createLinkedListUnsorted()
	fmt.Println("Original linked list:")
	printLinkedList(head)
	sortedHead := MergeSortLinkedList(head)
	fmt.Println("Sorted linked list:")
	printLinkedList(sortedHead)

	// ----Quicksort Algorithm----
	fmt.Println("\n--- QuickSort Algorithim has begun ---")
	quickSortArr := []int{5,3,8,1,3,9,57,4,28,13,2,142,7}
	QuickSort(quickSortArr, 0,len(quickSortArr)-1)
	fmt.Println(quickSortArr)

	BubbleSortNumsPrac := []int{9,8,7,6,5,4,11,12,34,123,565,3123,123,1236,64,7654,124,1235,643,3,2,1,67}
	BubbleSort2(BubbleSortNumsPrac)
	fmt.Println("Practice BubbleSort",BubbleSortNumsPrac)

	SelectionSortNumsPrac := []int{9,8,7,6,5,4,3,2,1,11,12,34,64,7654,123,1236,124,1235,643,67,123,565,3123}
	SelectionSortPrac(SelectionSortNumsPrac)
	fmt.Println("Selection Sort Practice",SelectionSortNumsPrac)

	InsertionSortNumsPrac := []int{9,8,7,6,5,4,3,2,1,11,12,34,64,7654,123,1236,124,1235,643,67,123,565,3123}
	InsertionSortPrac(InsertionSortNumsPrac)
	fmt.Println("Insertion Sort Practice",InsertionSortNumsPrac)

	MergeSortNumPrac := []int{9,8,7,6,5,4,3,2,1,11,12,34,64,7654,123,1236,124,1235,643,67,123,565,3123}
	sortedPrac := MergeSortPrac(MergeSortNumPrac)
	fmt.Println("Merge Sort Prac:",sortedPrac)

	partitionSortArr := []int{5,3,8,1,3,9,57,4,28,13,2,142,7}
	QuickSortPrac(partitionSortArr, 0 , len(partitionSortArr)-1)
	fmt.Println("QuickSort Practice",partitionSortArr)

}

func createLinkedListUnsorted()(*Node, *Node){
	MergeHead := &Node{val:0}
	MergeHead.next = &Node{val:1}
	MergeHead.next.next = &Node{val:5}
	MergeHead.next.next.next = &Node{val:10}
	MergeHead.next.next.next.next = &Node{val:-1}

	head := &Node{val: 5}
	head.next = &Node{val: 2}
	head.next.next = &Node{val: 1}
	head.next.next.next = &Node{val: 3}
	head.next.next.next.next = &Node{val: 4}
	return MergeHead, head

}

func printLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.val, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}


