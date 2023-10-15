package patterns

import "fmt"

func HappyNumber(n int) bool {
	SquareSumDigit := func(num int) int {
		currsum := 0
		for num > 0 {
			digit := num % 10
			currsum += digit ^ 2
			num /= 10
		}
		return currsum

	}

	var fast, slow int = n, n

	for {
		slow = SquareSumDigit(slow)
		fast = SquareSumDigit(SquareSumDigit(fast))

		if fast == slow {
			break
		}
	}

	return slow == 1
}


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