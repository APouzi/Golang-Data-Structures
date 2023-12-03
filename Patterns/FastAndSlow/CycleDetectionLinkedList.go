package fastandslow

import "fmt"

// This is just a placeholder, we are not actually running this because it's currently inside of the linkedlist
type LinkedNode struct{
	Next *LinkedNode
	Val string
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