package main

import "fmt"

func (ll *LinkedList) addNode(num int) {
	newHead := &LinkedNode{val: num}
	if ll.head == nil {
		ll.head = newHead
		return
	}
	curr := ll.head

	newHead.next = curr
	ll.head = newHead

}

func (ll *LinkedList) printList() {

	curr := ll.head
	for curr != nil {
		fmt.Printf(" %v ", curr.val)
		curr = curr.next
	}
}

func printReverse(node *LinkedNode) {
	curr := node

	for curr != nil {
		fmt.Printf(" %v ", curr.val)
		curr = curr.next
	}

}

func (ll *LinkedList) makeCycle() {
	curr := ll.head
	cycleConnect := curr.next.next
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = cycleConnect

}

func cycleDetect(node *LinkedNode) bool {
	fast := node
	slow := node

	for fast != nil || slow != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			fmt.Println("connected at: ", fast.val)
			return true
		}
	}
	return false
}

func revRecur(node *LinkedNode) *LinkedNode {
	if node.next == nil {

		return node
	}

	newHead := revRecur(node.next)
	node.next.next = node
	node.next = nil
	return newHead

}

//***************Linked List******************
type LinkedList struct {
	head *LinkedNode
}

type LinkedNode struct {
	val  int
	next *LinkedNode
}
