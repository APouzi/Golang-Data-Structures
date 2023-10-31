package sorting

import "fmt"

func BubbleSort(arr []int) {
	steps := 0
	for i := 0; i < len(arr); i ++{
		for j := 1; j < len(arr)-i;j++{
			steps++
			if arr[j] < arr[j-1]{
				arr[j], arr[j-1]= arr[j-1], arr[j]
			}
			
		}
	}
	fmt.Println("steps bubblesort 1:",steps)
}

func BubbleSort2(arr []int) {
	steps := 0
	for i := 0; i < len(arr); i ++{
		for j := 1; j < len(arr)-i-1;j++{
			steps++
			if arr[j] < arr[j+1]{
				arr[j], arr[j+1]= arr[j+1], arr[j]
			}
			
		}
	}
	fmt.Println("steps bubblesort :",steps)
}

func BubbleSort3(arr []int) {
	steps := 0
	for i := 0; i < len(arr); i ++{
		for j := i; j < len(arr);j++{
			steps++
			if arr[j] < arr[i]{
				arr[j], arr[i]= arr[i], arr[j]
			}
			
		}
	}
	fmt.Println("steps bubblesort :",steps)
}