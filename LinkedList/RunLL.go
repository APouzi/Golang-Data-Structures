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
	reved2 := ReverseKGroups(L3.Head, 3)
	println("\nprint reverse:")
	PrintAfterReverse(reved2)
	// newHead := revRecur(LinkedList.head)
	// RemoveDup(LL.Head)

	fmt.Println("\n\n---LinkedList Practice has started---")
	L4 := CreateLinkedList()
	fmt.Println("Reverse K Groups Practice Linked List:")
	L4.PrintList()
	revedPrac2 := ReverseKGroupsPrac(L4.Head, 3)
	println("\nprint reverse:")
	PrintAfterReverse(revedPrac2)

	L5 := CreateLinkedList()
	fmt.Println("\nLinkedList Middle Node:", MiddleNodeofLLPrac(L5.Head))

	L6 := CreateLinkedList()
	fmt.Println("LinkedList Reversal Recursive Prac:")
	PrintAfterReverse(RevLLRecurPrac(L6.Head))

	LCycle := CreateCycleLinkedList()
	fmt.Println("\nLinkedList Cycle:", CycleDetectionPrac(LCycle.Head))

	L7 := CreateLinkedList()
	fmt.Println("\nLinked List Reverse Iterative Practice:")
	PrintAfterReverse(RevLLIterativePrac(L7.Head))

}

func CreateLinkedList() *LinkedList{
	list := LinkedList{}
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	return &list
}

func CreateCycleLinkedList() *LinkedList{
	list := LinkedList{}
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	list.MakeCycle()
	return &list
}