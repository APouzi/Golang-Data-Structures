package LinkedList

// Given the head of a linked list and an integer val, remove all the nodes of the linked list that has Node.val == val, and return the new head.

// Input: head = [1,2,6,3,4,5,6], val = 6
// Output: [1,2,3,4,5]

// Input: head = [], val = 1
// Output: []

// Input: head = [7,7,7,7], val = 7
// Output: []
func RemoveElements(head *LinkedNode, val int) *LinkedNode {
    var dummyHead *LinkedNode = &LinkedNode{} //In this example, dummyhead helps us take care of things like starting with a full linkedList and then ending up with an empty list. See example "7,7,7,7"
    dummyHead.Next = head

    var curr *LinkedNode = dummyHead
	//We do curr != nil because of multiple reasons. 
	//1), We only enter this if there is a linkedNode to actually do work on, sometimes we can be given a nil.
	//2) In example of "1,2,6,3,4,5,6" if we delete the last 6, we will end up being a nil because of "curr.Next = curr.Next.Next" in the inner for loop. The problem is that if we try to ask the conditional of "curr.Next", we can't because .Next doesn't exist on nil. 
	//3) After our processing in the inner loop, there is a chance the linked list will become completely nil, so we need to stop work on it."7,7,7,7"
    for curr != nil{
        for curr.Next != nil  && curr.Next.Val == val{ //We need to ask curr.Next first because of the fact that if we can reach the beyond end the list and make curr.Next into nil. in example of "1,2,6,3,4,5,6", if we reach 5.Next, we will delete 6 and then .Next becomes nill, if we try doing 5.Next.Next, it will mean the 2nd .Next will be "attached" to nil and thus break.
            curr.Next = curr.Next.Next
        }
        curr = curr.Next
    }
    return dummyHead.Next
}