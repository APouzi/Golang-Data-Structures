package LinkedList

import "fmt"

func RunLinkedList() {
	LL := CreateLinkedList()
	L2 := CreateLinkedList()
	L3 := CreateLinkedList()
	
	reved := RevRecur(LL.Head)
	println("\nprint reverse:")
	PrintWithHead(reved)
	
	//---Remove Duplicated in Linked List
	fmt.Println("\nRemove Duplicate LinkedList")
	LLDup := CreateLinkedListDuplicate()
	LLDup.PrintList()
	fmt.Println("\nRemove Duplicate LinkedList Practice")
	deDupped := RemoveDupPrac(LLDup.Head)
	PrintWithHead(deDupped)	

	//---Reverse K Groups Linked List---
	fmt.Println("\nReverse K Groups Linked List:")
	L3.PrintList()
	reved2 := ReverseKGroups(L3.Head, 3)
	println("\nprint reverse:")
	PrintWithHead(reved2)
	fmt.Println("Reverse K Groups Practice Linked List:")
	L4 := CreateLinkedList()
	L4.PrintList()
	revedPrac2 := ReverseKGroupsPrac(L4.Head, 2)
	println("\nprint reverse:")
	PrintWithHead(revedPrac2)

	//---Middle Node of Linked List---
	L5 := CreateLinkedList()
	fmt.Println("\nLinkedList Middle Node:", MiddleNodeofLLPrac(L5.Head))

	//---Reverse Linked List Iterative---
	fmt.Println("\nReverse Linked List Iterative:")
	L2.PrintList()
	revedPrac := LLRevIterative(L2.Head)
	println("\nprint reverse:")
	PrintWithHead(revedPrac)
	L7 := CreateLinkedList()
	fmt.Println("\nLinked List Reverse Iterative Practice:")
	PrintWithHead(RevLLIterativePrac(L7.Head))

	//---LinkedList Reversal Recursive Version---
	fmt.Println("---LinkedList has started---")
	fmt.Println("\nReverse Linked List Recursive:")
	LL.PrintList()
	fmt.Println("LinkedList Reversal Recursive Prac:")
	L6 := CreateLinkedList()
	L6.PrintList()
	PrintWithHead(RevLLRecurPrac(L6.Head))

	//---Linked List Cycle Detection---
	LCycle := CreateCycleLinkedList()
	fmt.Println("\nLinkedList Cycle:", CycleDetectionPrac(LCycle.Head))

	

	//---Remove Nth Node from Linked List---
	L8 := CreateLinkedList()
	fmt.Println("\nRemove Nth Node from Linked List:")
	PrintWithHead(RemoveNthFromEnd(L8.Head, 3))
	fmt.Println("\nRemove Nth Node from Linked List Practice:")
	L9 := CreateLinkedList()
	PrintWithHead(RemoveNthNodePrac(L9.Head, 3))

	//-----Remove Element -----
	fmt.Println("Remove Elements from List:")
	L10 := CreateLinkedListRemoveElem()
	PrintWithHead(L10.Head)
	ret := RemoveElementsPrac(L10.Head, 7)
	PrintWithHead(ret)
	
	//----Odd Even LinkedList---
	L11 := CreateLinkedList()
	fmt.Println("\nOdd Even LinkedList:")
	oddeveRet := OddEvenList(L11.Head)
	PrintWithHead(oddeveRet)
	fmt.Println("\nOdd Even LinkedList Prac:")
	L12 := CreateLinkedList()
	oddeveRetPrac := OddEvenListPrac(L12.Head)
	PrintWithHead(oddeveRetPrac)

	// ----Merge 2 Sorted Lists---
	L13, L14 := MergeTwoLists(3) 
	mergedRet := MergeTwoSortedLists(L13,L14)
	fmt.Println("\nMerge Two Sorted Lists:")
	PrintWithHead(mergedRet)

	// ----Merge k Sorted Lists---
	lists := MergedLists(3,5) 
	mergedKRet := MergeKSortedLists(lists)
	fmt.Println("\nMerge K Sorted Lists:")
	PrintWithHead(mergedKRet)

}

func MergeTwoLists(listsNeeded int) (*LinkedNode, *LinkedNode){
	L1, L2 := &LinkedList{},&LinkedList{}
	for i := 10; i > 0; i--{
		if i %2 ==0{
			L1.AddNode(i)
		}else{
			L2.AddNode(i)
		}
	}
	return L1.Head, L2.Head
}

func MergedLists(listsNeeded int, countPerList int) []*LinkedNode{
	lists := []*LinkedNode{}
	k := 0
	for i := 0; i < listsNeeded; i++{
		lists = append(lists, &LinkedNode{Val: k})
		k+=countPerList
	}
	var count, start int
	for _, v := range lists{
		count = countPerList
		start = v.Val
		for count > 1 {
			start++
			v.Next = &LinkedNode{Val: start}
			v = v.Next
			count--
		}
	}
	fmt.Print("\nList of Linked Nodes Lists Input: ")
	for _, v := range lists{
		fmt.Print("[ ")
		for v != nil{
			fmt.Print(v.Val, " ")
			v = v.Next
		}
		fmt.Print("] ")
	}
	return lists
}


func CreateLinkedList() *LinkedList{
	list := LinkedList{}
	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}
	return &list
}

func CreateLinkedListRemoveElem() *LinkedList{
	list := LinkedList{}
	for i := 1; i < 10; i++ {
		if i == 7{
			for k := i; k < 10;k++{
				list.AddNode(i)
			}
			continue
		}
		list.AddNode(i)
	}
	list.AddNode(7)
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

func CreateLinkedListDuplicate()*LinkedList{
	list := LinkedList{}
	for i := 1; i < 10; i++ {
		if i == 5{
			list.AddNode(i)
			list.AddNode(i)
		}
		list.AddNode(i)
	}
	return &list
}

