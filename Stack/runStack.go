package stack

import "fmt"

func RunStackProbs() {
	fmt.Println("\nisValid started")
	// s := "()" //Pass
	s := "(())[]" //Pass
	// s := "){" //Fail
	// s:= "[(([))]]" //Fail
	fmt.Println(isValid(s))

	fmt.Println("Daily temperatures:", dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
}