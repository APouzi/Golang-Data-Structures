package prbGen

import (
	"fmt"
	"math/rand"
)
var AllProblems []string = []string{"mirror matrix", "rotate matrix", "spiral matrix", "where will the ball drop", "set matrix zeros", "Linked list", "middle of list", " reverse linked list iterative", "reverse k groups", "reverse linked list recursive", "delete duplicates linked list", "delete specific node","create a binary heap","merge intervals","insert interval","remove minimum over lapping interval","happy number", "valid parenthesis","DFS a Graph","BFS a Graph", "Union Find a Graph", "Number of Islands","Class Schedule", "valid palindrome", "valid anagram", "Sort a Linked List", "Quick Sort", "Merge Sort", "Binary Search", "Permutation of a String"}
func GenerationProblemOne() {
	
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
	fmt.Print("\n")
	for _,v := range problems{
		fmt.Println(v)
	}
}