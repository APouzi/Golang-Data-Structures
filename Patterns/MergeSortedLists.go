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






// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.

// Example 1:
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]

// Example 2:
// Input: list1 = [], list2 = []
// Output: []

// Example 3:
// Input: list1 = [], list2 = [0]
// Output: [0]

// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

//     if list1 == nil || list2 == nil {
//         if list1 == nil{
//             return list2
//         }
//         if list2 == nil{
//             return list1
//         }

//     }

//     minHeap := MinHeap{}
//     minHeap.push(list1)
//     minHeap.push(list2)
//     dummyHead := &ListNode{Val:-111}
//     mergedTail := dummyHead
//     for len(minHeap.array) > 0{
//         popped := minHeap.pop()
//         mergedTail.Next = popped
//         mergedTail = popped
//         if popped.Next != nil{
//             minHeap.push(popped.Next)
//         }
//     }
//     return dummyHead.Next
// }