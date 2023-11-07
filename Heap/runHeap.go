package heaps

import "fmt"

func RunHeap() {
	// LinkedList2 := LinkedList{}
	// LinkedList2.addNode(5)
	// LinkedList2.addNode(4)
	// LinkedList2.addNode(1)

	// LinkedList3 := LinkedList{}
	// LinkedList3.addNode(4)
	// LinkedList3.addNode(3)
	// LinkedList3.addNode(1)

	// LinkedList4 := LinkedList{}
	// LinkedList4.addNode(2)
	// LinkedList4.addNode(1)

	// kWayMergeList := []*LinkedNode{}
	// kWayMergeList = append(kWayMergeList, LinkedList2.head)
	// kWayMergeList = append(kWayMergeList, LinkedList3.head)
	// kWayMergeList = append(kWayMergeList, LinkedList4.head)
	// // printList(kWayMergeList)
	// // variable := kWayMerge(kWayMergeList)
	// binaryTree := BST{}
	// binaryTree.insertNode(5)
	// binaryTree.insertNode(3)
	// binaryTree.insertNode(4)
	// binaryTree.insertNode(8)
	// binaryTree.insertNode(6)
	// binaryTree.insertNode(7)
	// binaryTree.insertNode(9)
	// fmt.Println(kthbiggestElement(binaryTree.node, 5))

	// h := Heap{}
	// h.insert(3)

	// h.insert(1)

	// h.insert(7)
	// h.insert(8)
	// h.insert(9)
	// h.insert(4)
	// h.insert(5)
	// h.insert(6)
	// h.insert(2)
	// fmt.Println("final ans:", h.array)

	// for range h.array {
	// 	fmt.Println(h.pop())
	// }
	// fmt.Println(h.pop())
	// fmt.Println(h.pop())
	// fmt.Println(h.pop())
	// fmt.Println(h.pop())
	// fmt.Println(h.pop())
	// fmt.Println(h.pop())
	// fmt.Println(h.pop())
	// fmt.Println(h.pop())
	// fmt.Println(h.pop())
	// fmt.Println(variable)
	// printKWay(variable)

	fmt.Println("Started Heap Practice")
	insertIntoHeap := []int{3,1,4,10,11,2,10,12,88,2}
	heappractice := initHeapPrac()
	InsertandPopPrac(insertIntoHeap,heappractice)
}

func InsertandPopPrac(arr []int, heap *HeapPrac){
	for _,v := range arr{
		heap.InsertPrac(v)
	}

	for range heap.Heap{
		fmt.Println(heap.PopPrac())
	}
}