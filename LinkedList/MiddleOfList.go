package LinkedList

func MiddleOfList(head *LinkedNode) *LinkedNode {
	var slow, fast *LinkedNode
	//To understand why we need both, is because this accounts for both even and odd number of nodes. if it's 6 nodes, then fast will reach nil and slow will be in the middle. If it's 5, then fast will be at 5 and stop there, while slow being at nil.
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}