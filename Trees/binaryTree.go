package bst

import (
	"fmt"
)

func main() {
	binaryTree := BST{}
	binaryTree.insertNode(5)
	binaryTree.insertNode(3)
	binaryTree.insertNode(4)
	binaryTree.insertNode(8)
	binaryTree.insertNode(6)
	binaryTree.insertNode(7)
	binaryTree.insertNode(9)
	//		   5
	// 		3     4
	//               8
	//                  7
	//                     8
	//                        9

	binaryTreeValidate := BST{}
	binaryTreeValidate.node = &Node{val: 11}
	insertNodeValidate(binaryTreeValidate.node)
	traversal(binaryTreeValidate.node)
	// BFSonBST(binaryTree.node)
	// rightView(binaryTree.node)
	// fmt.Println("longest diameter:", longestDiameter(binaryTree.node))
	// ans := 0
	// bstLongest(binaryTree.node, &ans)
	// fmt.Println("longest diameter val:", ans)
	// fmt.Println(sumPath(binaryTree.node, 22, 0))
	// sums := []int{}
	// fmt.Println("biggest path", biggestPath(binaryTree.node, &sums, 0), sums)
	// listofLists := [][]int{}
	// checkSumRecur(binaryTree.node, 26, 0, []int{}, listofLists)
	fmt.Println("validate binary tree:", validateTreeInit(binaryTreeValidate.node))
	InOrderTraversal(binaryTree.node)

}

type BST struct {
	node *Node
}

type Node struct {
	val   int
	right *Node
	left  *Node
}

func insertNodeValidate(node *Node) {
	node.right = &Node{val: 14}
	node.left = &Node{val: 5}
	node.right.right = &Node{val: 15}
	node.right.left = &Node{val: 10}
	node.left.right = &Node{val: 8}
	node.left.left = &Node{val: 3}
	node.left.left.left = &Node{val: 1}
	node.left.left.right = &Node{val: 4}
	node.left.right.left = &Node{val: 7}
	node.left.right.right = &Node{val: 9}
}

func (bst *BST) insertNode(num int) {
	if bst.node == nil {
		bst.node = &Node{val: num}
		return
	}

	send := bst.node

	insertRecur(send, num)

}

func insertRecur(node *Node, num int) *Node {
	if node == nil {
		return &Node{val: num}
	}

	if num > node.val {
		node.right = insertRecur(node.right, num)
	}

	if num < node.val {
		node.left = insertRecur(node.left, num)
	}

	return node
}



func traversal(node *Node) {
	fmt.Println(node.val)
	if node.left != nil {
		traversal(node.left)

	}
	if node.right != nil {
		traversal(node.right)
	}
}






func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}








func rightSide(node *Node, min int, max int) bool {
	return true
}


