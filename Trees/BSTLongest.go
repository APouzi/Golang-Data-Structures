package bst

func bstLongest(node *Node, ans *int) int {
	if node == nil {
		return 0
	}

	left := bstLongest(node.left, ans)
	right := bstLongest(node.right, ans)
	left = max(left, 0)
	right = max(right, 0)
	*ans = max(*ans, left+right+node.val)
	return max(left, right) + node.val
}
