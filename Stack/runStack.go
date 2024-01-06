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
	fmt.Println("Daily temperatures Prac:", dailyTemperaturesPrac([]int{73,74,75,71,69,72,76,73}))

	//Needs to be 9:
	fmt.Println("Reverse Polish Notation:", ReversePolishNotation([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
}