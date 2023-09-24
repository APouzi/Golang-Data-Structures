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

func BFSonBST(node *Node) {
	queue := []*Node{}
	queue = append(queue, node)
	path := []int{}

	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		if popped.left != nil {
			queue = append(queue, popped.left)
		}

		if popped.right != nil {
			queue = append(queue, popped.right)
		}

		path = append(path, popped.val)

	}
	fmt.Print(path)
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

func rightView(node *Node) {
	queue := []*Node{}
	queue = append(queue, node)
	path := []int{}

	for len(queue) > 0 {
		rightSide := -1
		currLength := len(queue)
		for i := 0; i < currLength; i++ {
			popped := queue[0]
			queue = queue[1:]

			if popped.left != nil {
				queue = append(queue, popped.left)
			}

			if popped.right != nil {
				queue = append(queue, popped.right)
			}
			rightSide = popped.val
		}

		path = append(path, rightSide)

	}
	fmt.Print(path)
}

func longestDiameter(node *Node) int {
	if node == nil {
		return 0
	}

	left := longestDiameter(node.left)
	right := longestDiameter(node.right)

	return max(left, right) + 1
}

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

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

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

func checkSumRecur(node *Node, sum int, temp int, listTemp []int, list [][]int) bool {

	if node == nil {
		return false
	}
	listTemp = append(listTemp, node.val)

	if node.left == nil && node.right == nil && temp+node.val == sum {

		list = append(list, listTemp)
		fmt.Println("reached:", list)
		return true
	}
	checkSumRecur(node.left, sum, temp+node.val, listTemp, list)
	checkSumRecur(node.right, sum, temp+node.val, listTemp, list)

	return true
}

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

func rightSide(node *Node, min int, max int) bool {
	return true
}


func InOrderTraversal(node *Node){
	if node.left != nil{
		InOrderTraversal(node.left)
	}
	fmt.Println(node.val)
	if node.right != nil {
		InOrderTraversal(node.right)
	}

}

func PostOrderTraversal(node *Node){
	if node.right != nil{
		PostOrderTraversal(node)
	}
	if node.left != nil{
		PostOrderTraversal(node)
	}
	fmt.Println(node.val)
}

func PreOrderTraversal(node *Node){
	fmt.Println(node.val)
	if node.right != nil{
		PostOrderTraversal(node)
	}
	if node.left != nil{
		PostOrderTraversal(node)
	}
}