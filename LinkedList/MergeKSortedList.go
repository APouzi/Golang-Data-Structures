package LinkedList

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order. Merge all the linked-lists into one sorted linked-list and return it.
//lists[i] is sorted in ascending order.

// Example 1:
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted list:
// 1->1->2->3->4->4->5->6

// Example 2:
// Input: lists = []
// Output: []

// Example 3:
// Input: lists = [[]]
// Output: []
func MergeKSortedLists(lists []*LinkedNode) *LinkedNode{
	if lists == nil || len(lists) == 0{
		return nil
	}

	h := MinHeapNode{}

	for i := 0; i < len(lists);i++{
		if lists[i] != nil{ //This needs to check for the nil, due to the fact it's possible to get a nil item.
			h.Insert(lists[i])
		}
	}

	dummyHead := &LinkedNode{}
	ans := dummyHead

	for len(h.arr)>0{
		popped := h.Pop()
		ans.Next = popped
		ans = ans.Next
		if popped.Next != nil{
			h.Insert(popped.Next)
		}
	}
	return dummyHead.Next
}