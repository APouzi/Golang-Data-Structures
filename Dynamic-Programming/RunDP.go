package dynamicprogramming

import "fmt"

func RunDP() {
	fmt.Println("\nDynamic Programming has started")
	// ---Can Jump Problem---
	jumpTrue := []int{2, 3, 1, 1, 4}
	// jumpFalse := []int{3, 2, 1, 0, 4}
	canit := canJump(jumpTrue)
	fmt.Println("JumpGame:",canit)
}