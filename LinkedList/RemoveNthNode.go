package LinkedList

//Given the head of a linked list, remove the nth node from the end of the list and return its head.
//Input: head = [1,2,3,4,5], n = 2
//Output: [1,2,3,5]

//Input: head = [1], n = 1
//Output: []

//head = [1,2], n = 1
//[1]

//head = [1,2,3], n = 2
//[1,3]

func RemoveNthFromEnd(head *LinkedNode, n int) *LinkedNode {
	var dummyHead *LinkedNode = &LinkedNode{} //Fails on (head = [1,2], n =2), if dummyhead isn't used then we return a delete head, when the new list is 2.
	dummyHead.Next = head
	var forward, prevBeforeRemoval *LinkedNode = dummyHead, dummyHead //We need the "forward" to travel forward to let us count down the "n". This will signal us to let us know when we can flag a node for removal
	for forward.Next != nil {
		if n <= 0 {// This statement only activates after the n has reached 0, we will continuously move the previous until forward reaches the last node. Due to this, this will be on a lag by one node.
			prevBeforeRemoval = prevBeforeRemoval.Next //We want to put our marker on the node before the node we want to delete.
		}
		forward = forward.Next
		n--//This n right here is extremely important to place AFTER the check if n <= 0, if we do it too early or above the "n <= 0" this would mean that we would move the back pointer too early and causing is to remove the wrong node.
	}
	var theNodeToDelete *LinkedNode = prevBeforeRemoval.Next //Create the node that will be deleted, since we are behind it.
	prevBeforeRemoval.Next = theNodeToDelete.Next//Deletion
	return dummyHead.Next //The reason for this is due to the fact that if we want to return an empty list that is completely deleted, this is going to be very hard if we try to return head. In fact alot of this based is off the fact that head could be deleted. 
}