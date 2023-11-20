package patterns

//---min and max---
func min(a, b int) int{
	if a < b{
		return a
	}
	return b
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}


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


//---Quick Sort Intervals---
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

//---Merge Sort Intervals---

func MergeSortInterval(arr [][]int, st, end int) [][]int{
	if st >= end{
		return arr
	}

	mid := (st + end)/2

	MergeSortInterval(arr, st, mid)
	MergeSortInterval(arr, mid+1, end)
	return MergeInterval(arr, st, mid, end)

}

func MergeInterval(arr [][]int,st, mid, end int) [][]int{
	//remember, mid - st +1 because left side can hold 1 element in index 0, so "0-0" doesn't work when trying to get size
	var leftLim, rightLim int = mid-st+1, end - mid
	var left, right [][]int = make([][]int, leftLim), make([][]int, rightLim)
	var i int
	for i = 0; i < leftLim;i++{
		left[i] = arr[st+i]
	}

	for i = 0; i< rightLim;i++{
		right[i] = arr[mid+1+i]
	}

	var l,r,k int = 0,0,st

	for l < leftLim && r < rightLim{
		//
		if left[l][0] <= right[r][0]{
			arr[k] = left[l]
			l++
			k++
		}else{
			arr[k] = right[r]
			r++
			k++
		}
	}

	for l < leftLim{
		arr[k] = left[l]
		l++
		k++
	}

	for r < rightLim{
		arr[k] = right[r]
		r++
		k++
	}
	return arr
}



//Heap for LinkedLists


type MinHeap struct{
    array []*LinkedNode
}

func(h *MinHeap) push(num *LinkedNode){
    h.array = append(h.array, num)
    h.heapifyUp()
}

func(h *MinHeap) pop() *LinkedNode{
    popped := h.array[0]
    var length int = len(h.array) -1
    h.array[0] = h.array[length]
    h.array = h.array[:length]
    h.heapifyDown()
    return popped
}

func(h *MinHeap) heapifyUp(){
    var index int = len(h.array)-1
    for h.array[parent(index)].Val > h.array[index].Val{
        h.swap(parent(index), index)
    }
    
}
func(h *MinHeap) heapifyDown(){
    var index int = 0
    var length int = len(h.array)-1
    smallest := leftChild(index)
    for leftChild(index) < length{
        if rightChild(index) < length {
            if h.array[leftChild(index)].Val < h.array[rightChild(index)].Val{
                smallest = leftChild(index)
            }else{
                smallest = rightChild(index)
            }
            
            if h.array[index].Val > h.array[smallest].Val{
                h.swap(index, smallest)
                index = smallest
            }else{
                break
            }
        }else{
            if h.array[index].Val > h.array[leftChild(index)].Val{
                h.swap(index, leftChild(index))
                index = leftChild(index)
            }else{
                break
            }
        }
    }
}

func parent(num int) int{
    return num/2
}

func leftChild(num int) int{
    return (num + 1)/2
}

func rightChild(num int) int{
    return (num + 2)/2
}

func(h *MinHeap) swap(i int, j int){
    h.array[j],h.array[i] = h.array[i], h.array[j]
}


//------------------Heap for Pairs------------------
type HeapPair struct {
	arr [][]int
}

func InitializeHeapPair() *Heap {
	return &Heap{arr: []int{}}
}

func (h *HeapPair) Pop() []int {
	popped := h.arr[0]
	var lastIndex int = len(h.arr) - 1
	h.arr[0] = h.arr[lastIndex]
	h.arr = h.arr[:lastIndex]
	h.HeapifyDown(0)
	return popped
}

func (h *HeapPair) Insert(num []int) {
	h.arr = append(h.arr, num)
	var lastIndex int = len(h.arr) - 1
	h.HeapifyUp(lastIndex)
}

func (h *HeapPair) HeapifyUp(index int) {
	for h.arr[index][0] < h.arr[h.Parent(index)][0] {
		h.arr[index], h.arr[h.Parent(index)] = h.arr[h.Parent(index)], h.arr[index]
		index = h.Parent(index)
	}
}

func (h *HeapPair) HeapifyDown(index int) {
	var lastIndex int = len(h.arr) - 1
	var smallestChild int
	for h.leftChild(index) <= lastIndex {
		smallestChild = h.leftChild(index)
		if h.rightChild(index) <= lastIndex {
			if h.arr[smallestChild][0] > h.arr[h.rightChild(index)][0] {
				smallestChild = h.rightChild(index)
			}
			if h.arr[smallestChild][0] < h.arr[index][0] {
				h.arr[smallestChild], h.arr[index] = h.arr[index], h.arr[smallestChild]
				index = smallestChild
			} else {
				return
			}
		}
		if h.arr[smallestChild][0] < h.arr[index][0] {
			h.arr[smallestChild], h.arr[index] = h.arr[index], h.arr[smallestChild]
		}
		index = smallestChild

	}
}

func (h *HeapPair) Delete(num int) {
	for i, v := range h.arr {
		if v[0] == num {
			h.arr = append(h.arr[:i], h.arr[i:]...)
		}
	}
}

func (h *HeapPair) Parent(index int) int {
	return (index - 1) / 2
}

func (h *HeapPair) leftChild(index int) int {
	return (index * 2) + 1
}

func (h *HeapPair) rightChild(index int) int {
	return (index * 2) + 2
}