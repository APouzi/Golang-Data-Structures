package dynamicprogramming

import "fmt"

func RunDP() {
	// ---Can Jump Problem---
	jumpTrue := []int{2, 3, 1, 1, 4}
	// jumpFalse := []int{3, 2, 1, 0, 4}
	canit := canJump(jumpTrue)
	fmt.Println("canJump can jump",canit)
}