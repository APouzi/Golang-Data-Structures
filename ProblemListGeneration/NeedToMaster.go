package prbGen

import (
	"fmt"
	"math/rand"
)

var MasterThese []string = []string{
	"Trapping Rain Water", "Container With Most Water", "Two Sum II", "Longest Repeating Characters With Replacement", "Max Sliding Window", "Skyline Problem", "Minimum Interval from Query", "Meeting Room II", "Maximum Population", "Happy Number", "Find Duplicate Element In Array", "Return Linked List Cycle start (Floyds)", "Reverse K Groups", "Reverse Linked List II", "Remove Nth node", "Missing Number", "Is Valid Palindrome", "Valid Sudoku", "Product of Array except itself", "Group Anagrams", "Valid Anagram", "Rotate Image",
}

func MasterGenerationProblemOne() {
	fmt.Println("\nTotal Problems in Knowledge Bank is:", len(AllProblems))
	fmt.Println("\n", AllProblems[rand.Intn(len(AllProblems))])
}

func MasterGenerationProblemList() {
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