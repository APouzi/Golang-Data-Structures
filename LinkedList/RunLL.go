package LinkedList

import "fmt"

func RunLinkedList() {
	LL := LinkedList{}
	for i := 1; i < 10; i++ {
		LL.AddNode(i)
	}

	L2 := LinkedList{}
	for i := 1; i < 10; i++ {
		L2.AddNode(i)
	}
	fmt.Println("ReverseKGroups")
	LL.PrintList()
	reved := ReverseKGroups(LL.Head, 3)
	println("\n print reverse:")
	PrintAfterReverse(reved)


	// newHead := revRecur(LinkedList.head)
	// RemoveDup(LL.Head)

}