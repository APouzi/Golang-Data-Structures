package stack

import "fmt"

func RunStackProbs() {
	fmt.Println("\nisValid started")
	// s := "()" //Pass
	// s := "(())[]" //Pass
	s := "){" //Fail
	//s := ")" //Fail
	// s:= "[(([))]]" //Fail
	fmt.Println("Valid Parenthis:",isValid(s))
	fmt.Println("Valid Parenthis Prac:",isValidPrac(s))

	fmt.Println("\nDaily temperatures:", dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
	fmt.Println("Daily temperatures Prac:", dailyTemperaturesPrac([]int{73,74,75,71,69,72,76,73}))

	//Needs to be 9:
	fmt.Println("\nReverse Polish Notation:", ReversePolishNotation([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))
	fmt.Println("Reverse Polish Notation:", ReversePolishNotationPrac([]string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}))

	fmt.Println("\nGenerates Parenthesis:", generateParenthesis(3))
	fmt.Println("Generates Parenthesis Prac:", generateParenthesisPrac(3))

	var speed, position []int = []int{10,8,0,5,3}, []int{2,4,1,1,3}
	var target int = 12
	// var speed, position []int = []int{3}, []int{3}
	// var target int = 10
	// var speed, position []int = []int{0,2,4}, []int{4,2,1}
	// var target int = 100
	fmt.Println("\nCar Fleet:", carFleet(target,speed,position ))
	fmt.Println("Car Fleet Practice:", carFleetPrac(target,speed,position))

	fmt.Println("\nLargest Rectangle in a Histogram", largestRectangleArea([]int{2,1,5,6,2,3}))
	fmt.Println("Largest Rectangle in a Histogram", largestRectangleAreaPrac([]int{2,1,5,6,2,3}))
}