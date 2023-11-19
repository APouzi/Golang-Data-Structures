package patterns

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order. Merge all the linked-lists into one sorted linked-list and return it.

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

// func mergeKLists(lists []*LinkedNode) *LinkedNode {
// 	if lists == nil || len(lists) == 0 {
// 		return nil
// 	}

// 	var dummyHead *LinkedNode = &LinkedNode{Val: -111}
// 	var mergedTail *LinkedNode = dummyHead
// 	minHeap := Heap{}
// 	for _, v := range lists {
// 		if v != nil {
// 			minHeap.Insert(v)
// 		}
// 	}

// 	for len(minHeap.arr) > 0 {
// 		popped := minHeap.Pop()
// 		mergedTail.Next = popped
// 		mergedTail = popped
// 		if popped.Next != nil {
// 			minHeap.Insert(popped.Next)
// 		}
// 	}
// 	mergedTail.Next = nil
// 	return dummyHead.Next
// }