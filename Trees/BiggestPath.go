package bst

import "fmt"

func biggestPath(node *Node, sums *[]int, temp int) bool {
	if node == nil {
		return true
	}
	if node.right == nil && node.left == nil {
		*sums = append(*sums, temp+node.val)
		fmt.Println(sums)
		return true
	}
	//
	biggestPath(node.left, sums, temp+node.val)
	biggestPath(node.right, sums, temp+node.val)
	return true
}