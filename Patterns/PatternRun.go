package patterns

import "fmt"

func RunPattern() {
	// ---Permutation String---
	returnRes := PermutationOfString("dcda", "adc")
	fmt.Println("PermutationOfString",returnRes)
	returnpracRes := PermuStrPrac("dcda", "adc")
	fmt.Println("PermuStrPrac result",returnpracRes)

	// --- High Sell low
	returnRes2 := BuyHighSellLow([]int{6, -5, 4, 6, 1, 2, 5, 6, 5, 4, 6, 1, 2, 5, 6, 5, 4, 6, 1, 2, 11})
	fmt.Println(returnRes2)


	// ---Permutation String---
	HappyInput := 19
	HappyNum := HappyNumber(HappyInput)
	fmt.Printf("Is the number \"%d\", a happy num? Answer: %t\n", HappyInput,HappyNum)
	HappyNumPrac := HappyNumberPrac(HappyInput)
	fmt.Printf("Practice! Is the number \"%d\", a happy num? Answer: %t\n", HappyInput,HappyNumPrac)

	// ---Two Pointers --
	palindrone := "ABBAABBAABBA"
	isPalin := ValidPalindrome(palindrone)
	fmt.Println("Is",palindrone," a palindrone?:",isPalin)


	//---Merge Intervals---
	intervals := [][]int{{2,5},{1,3},{1,4},{9,15},{6,10}, {7,13}}
	fmt.Println("Merge Intervals:",MergeIntervals(intervals))

	intervalsPrac := [][]int{{2,5},{1,3},{1,4},{9,15},{6,10}, {7,13}}
	fmt.Println("Merge Intervals Practice:",MergeIntervalsPrac(intervalsPrac))

	//---Insert Intervals---
	intervalsIns := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	fmt.Println("Insert Intervals:",InsertIntervals(intervalsIns,[]int{4,8}))

	intervalsInsPrac := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	fmt.Println("Insert Intervals Prac:",InsertIntervalsPrac(intervalsInsPrac,[]int{4,8}))

	//---Non OverLapping Intervals---
	intervalsRem := [][]int{{1,2},{2,3},{3,4},{1,3}}
	fmt.Println("Non OverLapping Intervals:", nonOverLappingIntervals(intervalsRem))
}