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
	fmt.Println()
	curr := ll.Head
	for curr != nil {
		fmt.Printf(" %v ", curr.Val)
		curr = curr.Next
	}
}

func PrintWithHead(node *LinkedNode) {
	curr := node
	for curr != nil {
		fmt.Printf(" %v ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()

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

	for fast != nil || slow != nil { //Here we are stopping if the linked list is a not a cycle
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { //If it is a cycle we are gauranteed to stop because eventually this will become true
			fmt.Println("connected at: ", fast.Val)
			return true
		}
	}
	return false
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
