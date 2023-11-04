package LinkedList

func RevRecur(node *LinkedNode) *LinkedNode {
	if node.Next == nil {

		return node
	}

	newHead := RevRecur(node.Next)
	newHead.Next = node //Can also be node.Next.Next
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