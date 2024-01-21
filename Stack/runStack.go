package stack

import "fmt"

func RunStackProbs() {
	fmt.Println("\nisValid started")
	// s := "()" //Pass
	s := "(())[]" //Pass
	// s := "){" //Fail
	// s:= "[(([))]]" //Fail
	fmt.Println("Valid Parenthis:",isValid(s))
	fmt.Println("Valid Parenthis Prac:",isValidPrac(s))

	fmt.Println("Daily temperatures:", dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
	fmt.Println("Daily temperatures Prac:", dailyTemperaturesPrac([]int{73,74,75,71,69,72,76,73}))

	//Needs to be 9:
	fmt.Println("Reverse Polish Notation:", ReversePolishNotation([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
	fmt.Println("Reverse Polish Notation:", ReversePolishNotationPrac([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))

	fmt.Println("Generates Parenthesis:", generateParenthesis(3))
	fmt.Println("Generates Parenthesis Prac:", generateParenthesisPrac(3))

	fmt.Println("Car Fleet:", carFleet(12,[]int{10,8,0,5,3}, []int{2,4,1,1,3}))
	fmt.Println("Car Fleet Practice:", carFleetPrac(12,[]int{10,8,0,5,3}, []int{2,4,1,1,3}))

	fmt.Println("Largest Rectangle in a Histogram", largestRectangleArea([]int{2,1,5,6,2,3}))
	fmt.Println("Largest Rectangle in a Histogram", largestRectangleAreaPrac([]int{2,1,5,6,2,3}))
}