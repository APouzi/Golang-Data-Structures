package stack

import "fmt"

func RunStack() {
	fmt.Println("\nisValid started")
	// s := "()" //Pass
	s := "(())[]" //Pass
	// s := "){" //Fail
	// s:= "[(([))]]" //Fail
	fmt.Println(isValid(s))
}