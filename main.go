package main

import (
	"flag"
	"fmt"

	arrays "github.com/APouzi/go-algos/Arrays"
	dynamicprogramming "github.com/APouzi/go-algos/Dynamic-Programming"
	graph "github.com/APouzi/go-algos/Graph"
	heaps "github.com/APouzi/go-algos/Heap"
	"github.com/APouzi/go-algos/LinkedList"
	matrices "github.com/APouzi/go-algos/Matrices"
	patterns "github.com/APouzi/go-algos/Patterns"
	monotonicdeque "github.com/APouzi/go-algos/Patterns/MonotonicDeque"
	prbGen "github.com/APouzi/go-algos/ProblemListGeneration"
	randomlc "github.com/APouzi/go-algos/RandomLC"
	sorting "github.com/APouzi/go-algos/Sorting"
	stack "github.com/APouzi/go-algos/Stack"
	bst "github.com/APouzi/go-algos/Trees"
)

func main() {
	var cmnd string
	flag.StringVar(&cmnd, "command", " ","Command Init:")
	flag.Parse()

	// graph := &Graph{}
	// graph.addVertex(1)
	// graph.addVertex(2)
	// graph.addVertex(3)
	// graph.addVertex(4)
	// graph.addVertex(5)

	// graph.addEdges(1, 2)
	// graph.addEdges(2, 3)
	// graph.addEdges(3, 4)
	// graph.addEdges(4, 5)
	// graph.topoSort()

	if cmnd == "duplication"{
		LL := LinkedList.LinkedList{}
		LL.AddNode(1)
		LL.AddNode(2)
		LL.AddNode(3)
		LL.AddNode(3)
		LL.AddNode(3)
		LL.AddNode(4)
		LL.AddNode(5)
		LL.AddNode(6)

		LL.PrintList()
		println("\n print reverse:")
		LinkedList.RemoveDup(LL.Head)
		LL.PrintList()
	}

	// if cmnd == "reverse-recursive"{
	// 	LL := LinkedList.LinkedList{}
	// 	LL.AddNode(1)
	// 	LL.AddNode(2)
	// 	LL.AddNode(3)
	// 	LL.AddNode(4)
	// 	LL.AddNode(5)
	// 	LL.AddNode(6)

	// 	LL.PrintList()
	// 	reved := LinkedList.LLRev(LL.Head)
	// 	println("\n print reverse:")
	// 	LinkedList.PrintAfterReverse(reved)
	// 	// newHead := revRecur(LinkedList.head)
	// 	LinkedList.RemoveDup(LL.Head)
	// }
// "dcda", "adc"
// "eidbaooo", "eidbaooo"
// "abcdefg", "ab"
// "abcdefg", "dc"
	if cmnd == "pattern"{
		patterns.RunPattern()
	}

	if cmnd == "sort"{
		sorting.RunSorting()
	}

	if cmnd == "dynamic-programming"{
		dynamicprogramming.RunDP()
	}

	if cmnd == "stack"{
		stack.RunStackProbs()
	}

	if cmnd == "graph"{
		graph.Rungraphs()
	}

	if cmnd == "linked-list"{
		LinkedList.RunLinkedList()
	}

	if cmnd == "random-leet-code"{
		randomlc.RunRandomlC()
	}

	if cmnd == "heaps"{
		heaps.RunHeap()
	}

	if cmnd == "matrix"{
		matrices.RunMatrix()
	}

	if cmnd == "array"{
		arrays.RunArrays()
	}

	if cmnd == "tree"{
		bst.RunTrees()
	}

	if cmnd == "problem"{
		prbGen.GenerationProblemOne()
	}

	if cmnd == "problems"{
		prbGen.GenerationProblemList()
	}
	
	if cmnd == "mono"{
		monotonicdeque.Mono()
	}
	
	
	
	
	
	
	

	// fmt.Println(newHead.next)
	// printReverse(newHead)
	// LinkedList.printList()
	// LinkedList.makeCycle()
	// fmt.Println(cycleDetect(LinkedList.head))

	// fmt.Println("Two Baskets: ", twoBaskets([]int{1, 2, 2, 3, 1, 2, 2, 2, 1, 1, 2})) //7
	// BinaryTree := BST{}
	// BinaryTree.addNode(5)
	// BinaryTree.addNode(3)
	// BinaryTree.addNode(7)
	// BinaryTree.addNode(2)
	// BinaryTree.addNode(4)
	// BinaryTree.addNode(6)
	// BinaryTree.addNode(9)

	// printBST(BinaryTree.root)
	// // fmt.Println("Delete")
	// // deepCopy := []int{}
	// // arraysOfArrays := [][]int{}
	// // allPathsFromSource(BinaryTree.root, deepCopy, &arraysOfArrays, BinaryTree.root)
	// ans := 0
	// fmt.Println("maximum path: ", maximumOfPathSum(BinaryTree.root, &ans), ans)
	// // fmt.Println("All paths: ", arraysOfArrays)
	// // deleteNode(3, BinaryTree.root)
	// printBST(BinaryTree.root)
	// ans = 0
	// fmt.Println("The tree diameter is: ", treeDiam(BinaryTree.root, &ans), ans)
	// fmt.Println("The max path is: ", maxPathSum(BinaryTree.root))
	// rightSideView(BinaryTree.root)
	// BFS(BinaryTree.root)

}

// BinaryTree
type BST struct {
	root *BSTNode
}

type BSTNode struct {
	val   int
	right *BSTNode
	left  *BSTNode
}

//*****************Graph Structs***************
type Graph struct {
	vertices []*Node
}

type Node struct {
	val int
	adj map[*Node]*Node
}

// *****************************************************************Graph Functions***********************************************
func (g *Graph) addVertex(num int) {
	if g.vertices == nil {
		g.vertices = []*Node{}
	}

	new := &Node{val: num}
	new.adj = make(map[*Node]*Node)
	g.vertices = append(g.vertices, new)
}

func (g *Graph) addEdges(start int, end int) {
	var startNode *Node
	var endNode *Node
	for i := 0; i < len(g.vertices); i++ {
		if start == g.vertices[i].val {
			startNode = g.vertices[i]
		}
		if end == g.vertices[i].val {
			endNode = g.vertices[i]
		}
	}

	startNode.adj[endNode] = endNode
}

func (g *Graph) topoSort() {
	cycle := make(map[*Node]int)
	topoList := &[]int{}

	for _, v := range g.vertices {
		topological(v, cycle, topoList)
	}
	fmt.Println(topoList)
}

func topological(node *Node, cycle map[*Node]int, topoList *[]int) bool {
	if cycle[node] == 2 {
		return false
	}

	if cycle[node] == 1 {
		return false
	}

	cycle[node] = 2
	for _, v := range node.adj {
		if topological(v, cycle, topoList) == false {
			return false
		}
	}
	cycle[node] = 1
	*topoList = append(*topoList, node.val)

	return true
}

//*****************************************************Linked List Methods**********************************************

// two fruit baskets
func twoBaskets(fruits []int) int {
	bank := make(map[int]int)
	a := 0
	ans := 0

	for i := 0; i < len(fruits); i++ {
		for len(bank) > 2 {
			bank[fruits[a]] -= 1
			if bank[fruits[a]] == 0 {
				delete(bank, fruits[a])
			}

			a++
		}
		bank[fruits[i]] += 1
		ans = max(ans, i-a+1)
	}
	return ans

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// ---------------------------------------------Binary Search Tree Functions-------------------------------------------------

func addNode(num int, node *BSTNode) *BSTNode {
	if node == nil {
		return &BSTNode{val: num}

	}
	if node.val > num {
		node.left = addNode(num, node.left)
	} else if node.val < num {
		node.right = addNode(num, node.right)
	}
	return node
}

func deleteNode(num int, node *BSTNode) *BSTNode {
	if node == nil {
		return nil
	}
	if node.val > num {
		node.left = deleteNode(num, node.left)
	} else if node.val < num {
		node.right = deleteNode(num, node.right)
	} else {
		if node.right == nil {
			return node.left
		}
		if node.left == nil {
			return node.right
		}
		pred := findPredcessor(node.right)
		node.val = pred.val
		node.right = deleteNode(pred.val, node.right)
	}
	return node
}

func findPredcessor(node *BSTNode) *BSTNode {
	curr := node
	for curr.left != nil {
		curr = curr.left
	}
	return curr
}

func (bst *BST) addNode(num int) {
	if bst.root == nil {
		bst.root = &BSTNode{val: num}
	} else {
		addNode(num, bst.root)
	}
}

func printBST(node *BSTNode) {
	fmt.Println(node.val)
	if node.left != nil {
		printBST(node.left)
	}

	if node.right != nil {
		printBST(node.right)
	}

}

func maxPathSum(node *BSTNode) int {
	if node == nil {
		return 0
	}

	left := maxPathSum(node.left)
	right := maxPathSum(node.right)

	return max(left, right) + node.val
}

func treeDiam(node *BSTNode, ans *int) int {
	if node == nil {
		return 0
	}

	left := treeDiam(node.left, ans)
	right := treeDiam(node.right, ans)
	*ans = max(*ans, left+right+1)
	return max(left, right) + 1
}

func allPaths(node *BSTNode, deepCopy []int, shallowCopy *[][]int) int {
	if node == nil {
		// deepCopy = append(deepCopy, node.val)
		fmt.Println("deepCopy:", deepCopy)
		return 0
	}
	deepCopy = append(deepCopy, node.val)
	left := allPaths(node.left, deepCopy, shallowCopy)
	right := allPaths(node.right, deepCopy, shallowCopy)

	*shallowCopy = append(*shallowCopy, deepCopy)
	return max(left, right) + 1
}

func allPathsFromSource(node *BSTNode, deepCopy []int, shallowCopy *[][]int, source *BSTNode) {
	if node.right == nil && node.left == nil {
		deepCopy = append(deepCopy, node.val)
		*shallowCopy = append(*shallowCopy, deepCopy)
		return
	}

	deepCopy = append(deepCopy, node.val)
	allPathsFromSource(node.left, deepCopy, shallowCopy, source)
	allPathsFromSource(node.right, deepCopy, shallowCopy, source)
}

func maximumOfPathSum(node *BSTNode, ans *int) int {
	if node == nil {
		return 0
	}
	fmt.Print("node:", node.val)
	left := maximumOfPathSum(node.left, ans)
	right := maximumOfPathSum(node.right, ans)

	*ans = max(*ans, left+right+node.val)
	return max(left, right) + node.val
}

func rightSideView(node *BSTNode) []int {
	queue := []*BSTNode{}
	ans := []int{}
	queue = append(queue, node)
	for len(queue) > 0 {
		level := len(queue)
		rightSide := -101
		fmt.Println("queue length:", len(queue))
		for i := 0; i < level; i++ {

			popped := queue[0]
			queue = queue[1:]
			rightSide = popped.val
			if popped.left != nil {
				queue = append(queue, popped.left)
			}
			if popped.right != nil {
				queue = append(queue, popped.right)
			}

		}
		ans = append(ans, rightSide)
	}
	fmt.Println("rightSide:", ans)
	return ans
}

func BFS(node *BSTNode) []int {
	queue := []*BSTNode{}
	// seen := make(map[*BSTNode]bool)
	ans := []int{}
	queue = append(queue, node)
	for len(queue) > 0 {
		popped := queue[0]
		queue = queue[1:]
		if popped.left != nil {
			queue = append(queue, popped.left)
		}
		if popped.right != nil {
			queue = append(queue, popped.right)
		}
		ans = append(ans, popped.val)
	}
	fmt.Println(ans)
	return ans
}
