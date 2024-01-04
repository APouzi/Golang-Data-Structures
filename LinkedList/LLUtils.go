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
	var length int = len(h.arr) - 1
	h.arr[0] = h.arr[length]
	h.arr = h.arr[:length]
	h.HeapifyDown(0)
	return popped
}

func (h *MinHeapNode) HeapifyUp(num int) {
	for h.arr[h.Parent(num)].Val > h.arr[num].Val {
		h.Swap(h.Parent(num), num)
		num = h.Parent(num)
	}
}

func (h *MinHeapNode) HeapifyDown(num int) {
	var smallestChild int
	var length int = len(h.arr) - 1
	for h.LeftChild(num) <= length {
		smallestChild = h.LeftChild(num)
		if h.RightChild(num) <= length {
			if h.arr[h.RightChild(num)].Val < h.arr[smallestChild].Val {
				smallestChild = h.RightChild(num)
			}
			if h.arr[smallestChild].Val < h.arr[num].Val {
				h.Swap(smallestChild, num)
			} else {
				break
			}
		} else {
			if h.arr[smallestChild].Val < h.arr[num].Val {
				h.Swap(num, smallestChild)
			}
		}
		num = smallestChild
	}
}

func (h *MinHeapNode) Swap(a, b int) {
	h.arr[a], h.arr[b] = h.arr[b], h.arr[a]
}

func (h *MinHeapNode) Parent(num int) int {
	return (num - 1) / 2
}

func (h *MinHeapNode) LeftChild(num int) int {
	return (num * 2) + 1
}

func (h *MinHeapNode) RightChild(num int) int {
	return (num * 2) + 2
}