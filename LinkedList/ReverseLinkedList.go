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
	var curr *LinkedNode = currNode
	var prev *LinkedNode = nil
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}

// 1) loop through and put every kth node into the stack of the parameter of the head, make sure to build up the count, that will be the size of k
// 2) If the loop reaches nil of the list, we must return the nil, which will exit out the recursion and returnThis return will act as the previous to assign to as next
// When we peel off the stack, we will have access to the variables at each recursion.
// 3) As we peel back, this is where we reverse the linked nodes. At each peelback iteration, we do up to K, we do this by subtracting the incremented count.
// 4) return the newly assigned prev, which will be the current head, before we assign the head being the tempNext.
// 9  8  7  6  5  4  3  2  1	>>>		7  8  9 | 4  5  6 | 1  2  3
func ReverseKGroups(head *LinkedNode, k int) *LinkedNode {
	var copyCurr *LinkedNode = head
	var count int = 0
	for count < k { //This loop is going to be traverse until linked list reaches the end
		if copyCurr == nil { // When we return curr(3), this actually exit's out the recursion and start the peel back of the stack.
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