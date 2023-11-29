package prbGen

import (
	"fmt"
	"math/rand"
)
var AllProblems []string = []string{
	"mirror matrix", "rotate matrix", "spiral matrix", "where will the ball drop", "set matrix zeros", 
	"Linked list", "Cycle Detection","middle of list", " reverse linked list iterative", "reverse k groups", 
	"reverse linked list recursive", "delete duplicates linked list", "delete specific LL node","create a binary heap",
	"merge intervals","insert interval","remove minimum over lapping interval","happy number", "valid parenthesis",
	"DFS a Graph","BFS a Graph", "Union Find a Graph", "Number of Islands","Class Schedule", "valid palindrome", 
	"valid anagram", "Sort a Linked List", "Quick Sort", "Merge Sort", "Binary Search", "Permutation of a String", 
	"shuffle string", "Breadth First on BST", "BST Largest Sum Path", "Right View of Tree", "Longest Diameter of BST",
	"Right View of Tree", "Sum Path BST", "Validate BST", "In Order, Post Order, Pre Order Traversal of a Tree",
	"Check Sum BST", "Buy High Sell Low","Contains Duplicate", "Kth Smallest Element in a BST", "All Paths with Sum to 'k'",
	"Merge K Sorted Lists", "Longest Substring without repeating Characters", "Implement a Trie", "Fruits in a Basket",
	"Count all substring that equal to 'k", "Implement a Trie", "Keys and Rooms (Topological Sorting)", 
	"Sum of Nodes with Even-Valued Grandparent", "Course Schedule II", "Remove Nth Node From a List",
	"SubArray Product less than K","Merge K Sorted Lists","Merge 2 Sorted Lists", "Intersection of two Linked Lists",
	 "Maximum Subarray", "Delete Node from a BST", "Two Sum", "Path Sum II", "Invert a Binary Tree", "Meeting Rooms",
	"Meeting Rooms II", "Remove Elements", "skyline problem", "Top K frequent Elements", "Group Anagrams"}
func GenerationProblemOne() {
	fmt.Println("\nTotal Problems in Knowledge Bank is:",len(AllProblems))
	fmt.Println("\n",AllProblems[rand.Intn(len(AllProblems))])
}

func GenerationProblemList() {
	// indexes := make([]int,5)
	checkIfExists := map[int]bool{}
	count := 0
	problems := make([]string, 5)
	for count < 5{
		potentialNumber := rand.Intn(len(AllProblems))
		if _,ok := checkIfExists[potentialNumber]; !ok{
			problems[count] = AllProblems[potentialNumber]
			count++
		}
	}
	fmt.Println("\nTotal Problems in Knowledge Bank is:",len(AllProblems))
	for _,v := range problems{
		fmt.Println(v)
	}
}