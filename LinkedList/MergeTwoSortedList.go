package LinkedList

//You are given the heads of two sorted linked lists list1 and list2. Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.

//Example 1
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]

//Example 2:
// Input: list1 = [], list2 = []
// Output: []

//Example 3:
// Input: list1 = [], list2 = [0]
// Output: [0]

// Both list1 and list2 are sorted in non-decreasing order.

func MergeTwoSortedLists(L1, L2 *LinkedNode) *LinkedNode {
	if L1 == nil {
		return L2
	}
	if L2 == nil{
		return L1
	}
	h := MinHeapNode{}
	h.Insert(L1)
	h.Insert(L2)
	dummyHead := &LinkedNode{}
	ans := dummyHead
	for len(h.arr) > 0{
		popped := h.Pop()
		ans.Next = popped
		ans = ans.Next
		if popped.Next != nil{
			h.Insert(popped.Next)
		}
	}
	return dummyHead.Next
}