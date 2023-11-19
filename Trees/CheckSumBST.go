package bst

import "fmt"

func checkSumRecur(node *Node, sum int, temp int, listTemp []int, list [][]int) bool {

	if node == nil {
		return false
	}
	listTemp = append(listTemp, node.val)

	if node.left == nil && node.right == nil && temp+node.val == sum {

		list = append(list, listTemp)
		fmt.Println("reached:", list)
		return true
	}
	checkSumRecur(node.left, sum, temp+node.val, listTemp, list)
	checkSumRecur(node.right, sum, temp+node.val, listTemp, list)

	return true
}