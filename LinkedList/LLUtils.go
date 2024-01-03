package LinkedList

type MinHeapNode struct {
	arr []*LinkedNode
}

func (h *MinHeapNode) Insert(head *LinkedNode) {
	h.arr = append(h.arr, head)
	h.HeapifyUp(len(h.arr) - 1)
}

func (h *MinHeapNode) Pop() *LinkedNode {
	popped := h.arr[0]
	h.arr[0] = h.arr[len(h.arr)-1]
	h.arr = h.arr[:len(h.arr)-1]
	h.HeapifyDown(0)
	return popped
}

func (h *MinHeapNode) HeapifyUp(num int) {
	for h.Parent(num) >= 0 && h.arr[h.Parent(num)].Val > h.arr[num].Val {
		h.Swap(h.Parent(num), num)
		num = h.Parent(num)
	}
}

//	1
//
// 1 2
func (h *MinHeapNode) HeapifyDown(num int) {
	var smallestChild int
	for h.LeftChild(num) < len(h.arr)-1 {
		smallestChild = h.LeftChild(num)
		if h.RightChild(num) < len(h.arr)-1 && h.arr[smallestChild].Val > h.arr[h.RightChild(num)].Val {
			smallestChild = h.RightChild(num)
		}
		if h.arr[smallestChild].Val < h.arr[num].Val {
			h.Swap(num, smallestChild)
		} else {
			break
		}
		num = smallestChild
	}
	// if h.RightChild(num) < len(h.arr) && h.arr[h.RightChild(num)].Val < h.arr[num].Val {
	// 	h.Swap(h.RightChild(num), num)
	// }
}

func (h *MinHeapNode) Swap(a, b int) {
	h.arr[a], h.arr[b] = h.arr[b], h.arr[a]
}

// 1,2,3,4,5
func (h *MinHeapNode) Parent(num int) int {
	return (num - 1) / 2
}

func (h *MinHeapNode) LeftChild(num int) int {
	return (num * 2) + 1
}

func (h *MinHeapNode) RightChild(num int) int {
	return (num * 2) + 2
}