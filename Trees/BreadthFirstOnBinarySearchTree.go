package bst

import "fmt"

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