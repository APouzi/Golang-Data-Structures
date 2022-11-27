package main

import "fmt"

func main() {
	LinkedList2 := LinkedList{}
	LinkedList2.addNode(5)
	LinkedList2.addNode(4)
	LinkedList2.addNode(1)

	LinkedList3 := LinkedList{}
	LinkedList3.addNode(4)
	LinkedList3.addNode(3)
	LinkedList3.addNode(1)

	LinkedList4 := LinkedList{}
	LinkedList4.addNode(2)
	LinkedList4.addNode(1)

	kWayMergeList := []*LinkedNode{}
	kWayMergeList = append(kWayMergeList, LinkedList2.head)
	kWayMergeList = append(kWayMergeList, LinkedList3.head)
	kWayMergeList = append(kWayMergeList, LinkedList4.head)
	// printList(kWayMergeList)
	// variable := kWayMerge(kWayMergeList)
	binaryTree := BST{}
	binaryTree.insertNode(5)
	binaryTree.insertNode(3)
	binaryTree.insertNode(4)
	binaryTree.insertNode(8)
	binaryTree.insertNode(6)
	binaryTree.insertNode(7)
	binaryTree.insertNode(9)
	fmt.Println(kthbiggestElement(binaryTree.node, 5))

	h := Heap{}
	h.insert(3)
	h.insert(2)
	h.insert(1)

	h.insert(7)
	h.insert(8)
	h.insert(9)
	h.insert(4)
	h.insert(5)
	h.insert(6)
	fmt.Println(h.array)

	for range h.array {
		fmt.Println(h.pop())
	}

	// fmt.Println(variable)
	// printKWay(variable)

}

//*************************HEAP***************************
type LinkedList struct {
	head *LinkedNode
}

type LinkedNode struct {
	val  int
	next *LinkedNode
}

type Heap struct {
	array []int
}

func (h *Heap) insert(num int) {
	if h.array == nil {
		h.array = []int{}
	}
	// new := LinkedNode{val: node}
	h.array = append(h.array, num)
	h.heapifyUp()
}
func printKWay(node *LinkedNode) {
	for node.next != nil {
		fmt.Printf(" %v,", node.val)
		node = node.next
	}
}

func (h *Heap) heapifyUp() {
	index := len(h.array) - 1

	for h.array[index] < h.array[parent(index)] {
		h.swap(index, parent(index))
		index = parent(index)
	}
}

func (h *Heap) swap(a int, b int) {
	h.array[a], h.array[b] = h.array[b], h.array[a]
}

func (h *Heap) HeapifyDown() {
	length := len(h.array) - 1
	index := 0
	biggest := leftChild(index)
	for leftChild(index) < length {
		if rightChild(index) < length {
			if h.array[leftChild(index)] < h.array[rightChild(index)] {
				biggest = leftChild(index)
			} else {
				biggest = rightChild(index)
			}

			if h.array[index] > h.array[biggest] {
				h.swap(index, biggest)
				index = biggest
				// printList(h.array)
			} else {
				break
			}
		} else {
			if h.array[leftChild(index)] < h.array[index] {
				h.swap(leftChild(index), index)
				// printList(h.array)
			}
			index = leftChild(index)
		}
	}
}

func (h *Heap) pop() int {
	popped := h.array[0]
	length := len(h.array) - 1
	h.array[0] = h.array[length]
	h.array = h.array[:length]
	h.HeapifyDown()
	return popped
}

// 1,2,3,4,5,
func leftChild(index int) int {
	index = (index * 2) + 1
	return index
}

func parent(index int) int {
	index = (index - 1) / 2
	return index
}

func rightChild(index int) int {

	index = (index * 2) + 2
	return index
}

func printList(array []*LinkedNode) {
	fmt.Println()
	for _, v := range array {
		fmt.Printf(" %v ", v.val)
	}
}

// func kWayMerge(nodes []*LinkedNode) *LinkedNode {
// 	heap := Heap{}
// 	dummyHead := &LinkedNode{val: -1}
// 	tailHead := dummyHead
// 	for _, node := range nodes {
// 		heap.insert(node)

// 	}
// 	fmt.Println("heap length", len(heap.array))
// 	printList(heap.array)

// 	for len(heap.array) > 0 {
// 		popped := heap.pop()
// 		fmt.Println(popped.val)
// 		tailHead.next = &popped
// 		tailHead = &popped
// 		if popped.next != nil {
// 			heap.insert(popped.next)
// 		}
// 		// fmt.Println(popped.val)
// 	}
// 	return dummyHead.next
// }

// For KWay Merge
func (ll *LinkedList) addNode(num int) {
	newHead := &LinkedNode{val: num}
	if ll.head == nil {
		ll.head = newHead
		return
	}
	curr := ll.head

	newHead.next = curr
	ll.head = newHead

}

// kthbiggestElement

func kthbiggestElement(node *Node, k int) int {
	heap := Heap{}

	kthbiggestElementDFS(node, k, &heap)
	return heap.pop()
}

func kthbiggestElementDFS(node *Node, k int, heap *Heap) {
	if node == nil {
		return
	}

	heap.insert(node.val)
	if len(heap.array) > k {
		heap.pop()
	}
	kthbiggestElementDFS(node.left, k, heap)
	kthbiggestElementDFS(node.right, k, heap)
}

// BST
type BST struct {
	node *Node
}

type Node struct {
	val   int
	right *Node
	left  *Node
}

func (bst *BST) insertNode(num int) {
	if bst.node == nil {
		bst.node = &Node{val: num}
		return
	}

	send := bst.node

	insertRecur(send, num)

}

func insertRecur(node *Node, num int) *Node {
	if node == nil {
		return &Node{val: num}
	}

	if num > node.val {
		node.right = insertRecur(node.right, num)
	}

	if num < node.val {
		node.left = insertRecur(node.left, num)
	}

	return node
}
