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

	
}