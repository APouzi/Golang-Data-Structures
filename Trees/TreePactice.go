package bst

import "fmt"

func CreateTree() {
	//Create a tree
	//functions:
	//Insert Node
	//Traverse Node
	//Delete Node
	//Search Node
}
func biggestPathPrac(node *Node, sums *[]int, temp int) bool {
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

func validateTreePrac(node *Node, min int, max int) bool {
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

func InOrderTraversalPrac(node *Node) {
	if node.left != nil {
		InOrderTraversal(node.left)
	}
	fmt.Println(node.val)
	if node.right != nil {
		InOrderTraversal(node.right)
	}

}

func PostOrderTraversalPrac(node *Node) {
	if node.right != nil {
		PostOrderTraversal(node)
	}
	if node.left != nil {
		PostOrderTraversal(node)
	}
	fmt.Println(node.val)
}

func PreOrderTraversalPrac(node *Node) {
	fmt.Println(node.val)
	if node.right != nil {
		PostOrderTraversal(node)
	}
	if node.left != nil {
		PostOrderTraversal(node)
	}
}

// Given the root of a binary tree, return the sum of values of nodes with an even-valued grandparent. If there are no nodes with an even-valued grandparent, return 0.
// A grandparent of a node is the parent of its parent if it exists.

//Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
//Output: 18
//Explanation: The red nodes are the nodes with even-value grandparent while the blue nodes are the even-value grandparents.
// 	   6
// 	  /  \
// 	 7    8
//   /\    /\
//  2  7  1  3
// /  / \     \
//9   1  4      5
func SumofNodesWithEvenGrandParentsPrac(node *Node)int{
	return -1
}