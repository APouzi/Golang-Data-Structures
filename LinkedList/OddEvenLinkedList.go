package LinkedList

// Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

// The first node is considered odd, and the second node is even, and so on.

// Note that the relative order inside both the even and odd groups should remain as it was in the input.
// You must solve the problem in O(1) extra space complexity and O(n) time complexity.

// Example 1:
// Input: head = [1,2,3,4,5]
// Output: [1,3,5,2,4]

// Example 2:
// Input: head = [2,1,3,5,6,4,7]
// Output: [2,3,6,7,1,5,4]
 

func OddEvenList(head *LinkedNode) *LinkedNode {

	var dummyOddHead *LinkedNode = &LinkedNode{Val: -1}
	var dummyEvenHead *LinkedNode = &LinkedNode{Val: -1}
	dummyOddTail, dummyEventail := dummyOddHead, dummyEvenHead
	curr := head
	index := 0 
	for curr != nil{
		if index % 2 ==0{
			dummyEventail.Next = curr
			dummyEventail = curr
		}else{
			dummyOddTail.Next = curr
			dummyOddTail = curr
		}
		curr = curr.Next
		index++
	}
	dummyEventail.Next = dummyOddHead.Next
	dummyOddTail.Next = nil
    return dummyEvenHead.Next
}

func OddEvenListByValue(head *LinkedNode) *LinkedNode {

	var dummyOddHead *LinkedNode = &LinkedNode{Val: -1}
	var dummyEvenHead *LinkedNode = &LinkedNode{Val: -1}
	dummyOddTail, dummyEventail := dummyOddHead, dummyEvenHead
	curr := head

	for curr != nil{
		if curr.Val % 2 ==0{
			dummyEventail.Next = curr
			dummyEventail = curr
		}else{
			dummyOddTail.Next = curr
			dummyOddTail = curr
		}
		curr = curr.Next
	}
	dummyEventail.Next = dummyOddHead.Next
	dummyOddTail.Next = nil
    return dummyEvenHead.Next
}