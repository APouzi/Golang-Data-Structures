package bst

//Given a binary tree, find its minimum depth.

// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

// Note: A leaf is a node with no children.

//Example 1:
//Input: root = [3,9,20,null,null,15,7]
// Output: 2

// Example 2:
// Input: root = [2,null,3,null,4,null,5,null,6]
// Output: 5


func MinDepth(root *Node) int {
	if root == nil {
		return 0
	}

	return DfsMinDepth(root)
}

func DfsMinDepth(root *Node) int {
	if root == nil {
		return 0
	}

	var left int = DfsMinDepth(root.left)
	var right int = DfsMinDepth(root.right)

	if root.left == nil || root.right == nil { // This will get hit if we have an incomplete node, as in no left or right node. This is because we know that one of them will be zero, so, it won't matter that we are adding both because only one of them will have the depth. 
	//Also, note, if we have an imbalanced tree, as in just left or right nodes creating a straight line. Then this is the only one that will be hit, the "return min(left, right) + 1" will never be hit.
		return left + right + 1
	}
	//This one only gets hit at the root node(top of tree) or nodes that are complete and have both left and right nodes.
	return min(left, right) + 1
}


//     1
//    / \
//   2   3
//      /
//     4