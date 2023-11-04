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

func LLRev(currNode *LinkedNode) *LinkedNode { // 6,5,4,3,2,1
	// 6,5,4,3,   1 > 2 > nil
	// 6,5,4,,   1 > 2 > nil

	if currNode.Next == nil {
		return currNode
	}
	returnedHead := LLRev(currNode.Next)
	currNode.Next.Next = currNode
	currNode.Next = nil
	return returnedHead

}

// 1,2,3,4,5,6 > 3,2,1,4,5,6
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

// 9  8  7  6  5  4  3  2  1  to 7  8  9 | 4  5  6 | 1  2  3
func ReverseKGroups(curr *LinkedNode, k int) *LinkedNode {
	var copyCurr *LinkedNode = curr
	var count int = 0
	for count < k {
		if copyCurr == nil {
			return curr
		}
		copyCurr = copyCurr.Next
		count++
	}
	var prev *LinkedNode = ReverseKGroups(copyCurr, k)

	for count > 0 {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		count--
	}
	return prev
}