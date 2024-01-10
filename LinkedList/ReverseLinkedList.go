package LinkedList

func RevRecur(node *LinkedNode) *LinkedNode {
	if node.Next == nil {

		return node
	}

	newHead := RevRecur(node.Next)
	node.Next.Next = node //Can also be node.Next.Next
	node.Next = nil
	return newHead

}

// 9, 8, 7, 6, 5, 4, 3, 2, 1
func LLRevIterative(currNode *LinkedNode) *LinkedNode {
	var dummyHead *LinkedNode = &LinkedNode{}
	var curr *LinkedNode = currNode
	dummyHead.Next = curr
	var prev *LinkedNode = nil
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}



// k=3, 1,2,3,4,5,6 > 3,2,1,4,5,6
func LLRevKSendOff(curr *LinkedNode, k int) *LinkedNode {
	sendBack, connectTo, connect := LLRevK(curr, k, 1)
	connect.Next = connectTo
	return sendBack
}

func LLRevK(curr *LinkedNode, k, flag int) (*LinkedNode, *LinkedNode, *LinkedNode) {
	if flag == k || curr.Next == nil {
		return curr, curr.Next, nil
	}
	conn, retNext, _ := LLRevK(curr.Next, k, flag+1)
	curr.Next.Next = curr
	curr.Next = nil
	return conn, retNext, curr
}