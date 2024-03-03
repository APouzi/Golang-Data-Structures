package fastandslow

import "fmt"

func RunFastAndSlow() {
	// ---Happy Number---
	HappyInput := 19
	HappyNum := HappyNumber(HappyInput)
	fmt.Printf("\nIs the number \"%d\", a happy num? Answer: %t\n", HappyInput, HappyNum)
	HappyNumPrac := HappyNumberPrac(HappyInput)
	fmt.Printf("Practice! Is the number \"%d\", a happy num? Answer: %t\n", HappyInput, HappyNumPrac)
}