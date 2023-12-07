package monotonicdeque

import "fmt"

func Mono() {
	heights := []int{2, 1, 2, 4, 3}

	fmt.Println("Monotonic example heights", MonotonicArrayInit(heights))
}