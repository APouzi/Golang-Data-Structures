package LinkedList

func RemoveNthFromEnd(head *LinkedNode, n int) *LinkedNode {
	var dummyHead *LinkedNode = &LinkedNode{}
	dummyHead.Next = head
	var forward, prevBeforeRemoval *LinkedNode = dummyHead, dummyHead
	for forward.Next != nil {
		if n <= 0 {
			prevBeforeRemoval = prevBeforeRemoval.Next
		}
		forward = forward.Next
		n--
	}
	var theNodeToDelete *LinkedNode = prevBeforeRemoval.Next
	prevBeforeRemoval.Next = theNodeToDelete.Next
	return dummyHead.Next
}