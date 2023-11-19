package bst

import "fmt"

func InOrderTraversal(node *Node) {
	if node.left != nil {
		InOrderTraversal(node.left)
	}
	fmt.Println(node.val)
	if node.right != nil {
		InOrderTraversal(node.right)
	}

}

func PostOrderTraversal(node *Node) {
	if node.right != nil {
		PostOrderTraversal(node)
	}
	if node.left != nil {
		PostOrderTraversal(node)
	}
	fmt.Println(node.val)
}

func PreOrderTraversal(node *Node) {
	fmt.Println(node.val)
	if node.right != nil {
		PostOrderTraversal(node)
	}
	if node.left != nil {
		PostOrderTraversal(node)
	}
}