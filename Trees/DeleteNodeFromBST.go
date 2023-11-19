package bst

func DeleteNode(n *Node, key int) *Node {
	if n == nil {
		return n
	}

	if n.val > key {
		n.left = DeleteNode(n.left, key)
	} else if n.val < key {
		n.right = DeleteNode(n.right, key)
	} else if n.val == key {
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}

		var succ int = smallInright(n.right)
		n.val = succ
		n.right = DeleteNode(n.right, succ)
	}
	return n
}

func smallInright(n *Node) int {
	var curr *Node = n
	for curr.left != nil {
		curr = curr.left

	}
	return curr.val
}