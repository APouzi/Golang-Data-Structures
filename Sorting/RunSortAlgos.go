package sorting

import "fmt"

type Node struct{
	val int
	next *Node
}



func RunSorting() {
	BubbleSortNums := []int{9,8,7,6,5,4,3,2,1,11,12,34,64,7654,123,1236,124,1235,643,67,123,565,3123}
	InsertionSortNums := []int{9,8,7,6,5,4,3,2,1,11,12,34,64,7654,123,1236,124,1235,643,67,123,565,3123}
	// ----BubbleSort----
	BubbleSort(BubbleSortNums)
	fmt.Println("BubbleSort",BubbleSortNums)

	// ----Selection Sort----
	SelectionSortAlgo(InsertionSortNums)
	fmt.Println("Insertion Sort",InsertionSortNums)

	// ----------Merge Sort (Array and Linked List)----------
	fmt.Println("---Merge Sort Algorithm has begun---")
	// MergeSortNum := []int{9,8,7,6,5,4,3,2,1,11,12,34,64,7654,123,1236,124,1235,643,67,123,565,3123}
	MergeSortNum := []int{6,5,12,10,9,1}
	MergeSort(MergeSortNum, 0, len(MergeSortNum)-1)
	// sortedPrac := MergeSortPrac(MergeSortNum)
	fmt.Println("Merge Sort Algorithm",MergeSortNum)

	fmt.Println("---Merge Sort Algorithm 2 has begun---")
	MergeSortNum2 := []int{6,1,7,8,10}
	SortedMergeSort2 := MergeSort2(MergeSortNum2)
	fmt.Println("Merge Sort Algorithm", SortedMergeSort2)

	fmt.Println("--- Merge Sort Algorithm LinkedList has begun ---")
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

	fmt.Println("Original linked list:")
	printLinkedList(head)

	sortedHead := MergeSortLinkedList(head)

	fmt.Println("Sorted linked list:")
	printLinkedList(sortedHead)
}

func printLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.val, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}


