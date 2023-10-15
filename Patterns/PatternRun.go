package patterns

import "fmt"

func RunPattern() {
	returnRes := PermutationOfString("dcda", "adc")
	fmt.Println(returnRes)

	returnRes2 := BuyHighSellLow([]int{6, -5, 4, 6, 1, 2, 5, 6, 5, 4, 6, 1, 2, 5, 6, 5, 4, 6, 1, 2, 11})
	fmt.Println(returnRes2)

	HappyInput := 19
	HappyNum := HappyNumber(HappyInput)
	fmt.Printf("Is the number \"%d\", a happy num? Answer: %t\n", HappyInput,HappyNum)
}