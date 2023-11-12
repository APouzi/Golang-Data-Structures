package patterns

type Heap struct {
	arr []int
}

func InitializeHeap() *Heap {
	return &Heap{arr: []int{}}
}

func (h *Heap) Pop() int {
	popped := h.arr[0]
	var lastIndex int = len(h.arr) - 1
	h.arr[0] = h.arr[lastIndex]
	h.arr = h.arr[:lastIndex]
	h.HeapifyDown(0)
	return popped
}

func (h *Heap) Insert(num int) {
	h.arr = append(h.arr, num)
	var lastIndex int = len(h.arr) - 1
	h.HeapifyUp(lastIndex)
}

func (h *Heap) HeapifyUp(index int) {
	for h.arr[index] > h.arr[h.Parent(index)] {
		h.arr[index], h.arr[h.Parent(index)] = h.arr[h.Parent(index)], h.arr[index]
		index = h.Parent(index)
	}
}

func (h *Heap) HeapifyDown(index int) {
	var lastIndex int = len(h.arr) - 1
	var biggestChild int
	for h.leftChild(index) <= lastIndex {
		biggestChild = h.leftChild(index)
		if h.rightChild(index) <= lastIndex {
			if h.arr[biggestChild] < h.arr[h.rightChild(index)] {
				biggestChild = h.rightChild(index)
			}
			if h.arr[biggestChild] > h.arr[index] {
				h.arr[biggestChild], h.arr[index] = h.arr[index], h.arr[biggestChild]
				index = biggestChild
			} else {
				return
			}
		}
		if h.arr[biggestChild] > h.arr[index] {
			h.arr[biggestChild], h.arr[index] = h.arr[index], h.arr[biggestChild]
		}
		index = biggestChild

	}
}

func (h *Heap) Parent(index int) int {
	return (index - 1) / 2
}

func (h *Heap) leftChild(index int) int {
	return (index * 2) + 1
}

func (h *Heap) rightChild(index int) int {
	return (index * 2) + 2
}