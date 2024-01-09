package LinkedList

func RevLinkedListII(head *LinkedNode, left, right int) *LinkedNode {
	var dummyHead *LinkedNode = &LinkedNode{}
	dummyHead.Next = head
	var prevLeft, rightNode *LinkedNode = dummyHead, head
	for i := 0; i < left-1; i++ {
		prevLeft = rightNode
		rightNode = rightNode.Next
	}
	var prev *LinkedNode = nil
	for i := 0; i < right-left+1; i++ {
		nextRight := rightNode.Next
		rightNode.Next = prev
		prev = rightNode
		rightNode = nextRight
	}
	prevLeft.Next.Next = rightNode
	prevLeft.Next = prev
	return dummyHead.Next
}