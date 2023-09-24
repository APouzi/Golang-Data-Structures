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

	fast := node
	slow := node

	for fast != nil || slow != nil {
		if fast == slow {
			return true
		}
	}
	return false
}


		return node
	}

	return newHead

}

//***************Linked List******************
type LinkedList struct {
}

type LinkedNode struct {
}
