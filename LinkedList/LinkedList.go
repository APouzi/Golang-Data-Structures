package LinkedList

import "fmt"

func (ll *LinkedList) AddNode(num int) {
	newHead := &LinkedNode{Val: num}
	if ll.Head == nil {
		ll.Head = newHead
		return
	}
	curr := ll.Head

	newHead.Next = curr
	ll.Head = newHead

}

func (ll *LinkedList) PrintList() {

	curr := ll.Head
	for curr != nil {
		fmt.Printf(" %v ", curr.Val)
		curr = curr.Next
	}
}

func PrintAfterReverse(node *LinkedNode) {
	curr := node

	for curr != nil {
		fmt.Printf(" %v ", curr.Val)
		curr = curr.Next
	}

}

func (ll *LinkedList) MakeCycle() {
	curr := ll.Head
	cycleConnect := curr.Next.Next
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = cycleConnect

}

func CycleDetect(node *LinkedNode) bool {
	fast := node
	slow := node

	for fast != nil || slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fmt.Println("connected at: ", fast.Val)
			return true
		}
	}
	return false
}

func RevRecur(node *LinkedNode) *LinkedNode {
	if node.Next == nil {

		return node
	}

	newHead := RevRecur(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead

}

func LLRev(currNode *LinkedNode) *LinkedNode{ // 6,5,4,3,2,1 
											  // 6,5,4,3,   1 > 2 > nil
											  // 6,5,4,,   1 > 2 > nil

	if currNode.Next == nil{
		return currNode
	}
	returnedHead := LLRev(currNode.Next)
	currNode.Next.Next = currNode
	currNode.Next = nil
	return returnedHead
	
}


func RemoveDup(currNode *LinkedNode) {
	curr := currNode

	for curr.Next != nil{
		for curr.Val == curr.Next.Val{
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}
}


//***************Linked List******************
type LinkedList struct {
	Head *LinkedNode
}

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}
