package bst

import "fmt"

func sumPath(node *Node, sum int, temp int) bool {
	if node == nil {
		return false
	}
	if node.right == nil && node.left == nil && temp+node.val == sum {
		return true
	}
	fmt.Println(node.val)
	left := sumPath(node.left, sum, temp+node.val)
	right := sumPath(node.right, sum, temp+node.val)

	return left || right
}