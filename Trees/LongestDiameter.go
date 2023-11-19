package bst

func longestDiameter(node *Node) int {
	if node == nil {
		return 0
	}

	left := longestDiameter(node.left)
	right := longestDiameter(node.right)

	return max(left, right) + 1
}