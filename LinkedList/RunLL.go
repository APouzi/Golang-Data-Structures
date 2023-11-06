package LinkedList

import "fmt"

func RunLinkedList() {
	LL := CreateLinkedList()
	L2 := CreateLinkedList()

	fmt.Println("ReverseKGroups")
	LL.PrintList()
	reved := ReverseKGroups(LL.Head, 3)
	println("\n print reverse:")
	PrintAfterReverse(reved)
	
	fmt.Println("\n")
	fmt.Println(L2.Head.Val)
	revedPrac := LLRevIterative(L2.Head)
	println("\n print reverse Prac:")
	PrintAfterReverse(revedPrac)
	// newHead := revRecur(LinkedList.head)
	// RemoveDup(LL.Head)

}

func CreateLinkedList() *LinkedList{
	list := LinkedList{}
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	return &list
}