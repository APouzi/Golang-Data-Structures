package bst
func validateTreeInit(node *Node) bool {
	min := -9999999
	max := 9999999

	return validateTree(node, min, max)
}

func validateTree(node *Node, min int, max int) bool {
	if node == nil {
		return true
	}

	left := validateTree(node.left, min, node.val)
	right := validateTree(node.right, node.val, max)
	if node.val >= max || node.val <= min {
		return false
	}
	return left && right
}
