package LinkedList

import "fmt"

func RunLinkedList() {
	LL := CreateLinkedList()
	L2 := CreateLinkedList()
	L3 := CreateLinkedList()
	fmt.Println("---LinkedList has started---")
	fmt.Println("\nReverse Linked List Recursive:")
	LL.PrintList()
	reved := RevRecur(LL.Head)
	println("\nprint reverse:")
	PrintAfterReverse(reved)
	
	fmt.Println("\nReverse Linked List Iterative:")
	L2.PrintList()
	revedPrac := LLRevIterative(L2.Head)
	println("\nprint reverse:")
	PrintAfterReverse(revedPrac)

	fmt.Println("\nReverse K Groups Linked List:")
	L3.PrintList()
	revedPrac2 := ReverseKGroupsPrac(L3.Head, 3)
	println("\nprint reverse:")
	PrintAfterReverse(revedPrac2)
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