package bst

import "fmt"

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