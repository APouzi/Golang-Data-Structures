package patterns

//---Heap----
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

func (h *Heap) Delete(num int) {
	for i, v := range h.arr {
		if v == num {
			h.arr = append(h.arr[:i], h.arr[i:]...)
		}
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


//---Sort Intervals---
//CANNOT USE QUICKSORT TO SORT INTERVALS, UNSTABLE
func QuickSortIntervals(arr [][]int, low, high int) [][]int{
	if low >= high{
		return arr
	}
	pivot := PartionIntervalList(arr, low, high)
	QuickSortIntervals(arr, low, pivot-1)
	QuickSortIntervals(arr, pivot+1, high)
	return arr
}

func PartionIntervalList(arr [][]int, low, high int ) int{
	var pivot int = high
	var smallest int = low
	for i := 0; i < high; i++{
		if arr[pivot][0] >= arr[smallest][0]{
			arr[smallest], arr[i] = arr[i], arr[smallest]
			smallest++
		}
	}
	arr[smallest], arr[pivot] = arr[pivot], arr[smallest]
	return smallest
	
}