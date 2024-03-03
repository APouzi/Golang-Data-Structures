package slidingwindow

import "fmt"

func RunSlidingWindowRun() {
	// --- High Sell low
	returnRes2 := BuyHighSellLow([]int{6, -5, 4, 6, 1, 2, 5, 6, 5, 4, 6, 1, 2, 5, 6, 5, 4, 6, 1, 2, 11})
	returnRes3 := BuyHighSellLowPrac([]int{7,6,4,3,1})
	fmt.Println("Buy High Sell Low:",returnRes2)
	fmt.Println("Buy High Sell Low Prac:", returnRes3)


	//---Fruit Basket---
	fmt.Println("\nFruit Basket Practice", TotalFruitPrac([]int{1, 2, 3, 2, 2}))

	//---Character Replacement---
	fmt.Println("\nLongest Repeating Characters With Replacement Characters:", CharacterReplacement("AABABBA", 1))
	fmt.Println("Longest Repeating Characters With Replacement Characters Prac:", CharacterReplacementPrac("AABABBA", 1))

	//---Minimum Window SubString---
	fmt.Println("\nMinimum Window SubString Practice", MinWindowPrac("ADOBECODEBANC", "ABC"))

	//---Maximum Sliding Window ---
	fmt.Println("\nMaximum Sliding Window", MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println("Maximum Sliding Window Practice", MaxSlidingWindowPrac([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

	//---Maximum Product Subarray---
	fmt.Println("\nMaximum Product Subarray:", MaxProduct([]int{7, -2, -4}))
	fmt.Println("Maximum Product Subarray Prac:", MaxProductPrac([]int{7, -2, -4}))

	//---Longest SubString without Repeating Characters---
	fmt.Println("\nLongest Substring without Repeating Characters Practice", LengthOfLongestSubstringPrac("pwwkew"))

	// ---Permutation String---
	returnRes := PermutationOfString("eidboaoo", "eidboaoo")
	fmt.Println("\nPermutationOfString",returnRes)
	//"eidboaoo", str2 "eidboaoo"
	returnpracRes := PermutationOfStringPrac("eidboaoo", "ba")
	fmt.Println("PermuStrPrac result",returnpracRes)
}