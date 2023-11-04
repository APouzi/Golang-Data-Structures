package LinkedList

import "fmt"

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

// 9  8  7  6  5  4  3  2  1	>>>		7  8  9 | 4  5  6 | 1  2  3
func ReverseKGroups(head *LinkedNode, k int) *LinkedNode {
	var copyCurr *LinkedNode = head
	var count int = 0
	for count < k { //This loop is going to be traverse until linked list reaches the end
		if copyCurr == nil { // When we return curr(3), this actually exit's out the recursion and start the peel back of the stack.
			fmt.Println(head)
			return head //With the LL above, this will return "3".
		}
		copyCurr = copyCurr.Next
		count++
	}
	var last *LinkedNode = ReverseKGroups(copyCurr, k) //The first return will be 3
	// last = nil, last = 1, last = 4
	// head = 3, head = 6, head = 9
	for count > 0 { 
		next := head.Next // 2, 1, nil
		head.Next = last  // 3 > nil, 2 > 3, 1 > 2
		last = head       // nil = 3, 3 = 2, 2 = 1
		head = next       // 3 = 2, 2 = 1, 1 = nil
		count--
	}
	return last
}