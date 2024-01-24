package bst

import "fmt"

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

func hasPathSum(root *Node, targetSum int) bool {
    if root == nil{
        return false
    }
    return dfs(root, targetSum)
}

func dfs(root *Node, targetSum int) bool{
    if root == nil {
        return false
    }
	
	if root != nil && root.right == nil && root.left == nil && targetSum - root.val == 0{
        return true
    }
	

    return dfs(root.left, targetSum - root.val) || dfs(root.right, targetSum - root.val) //Since we are returning a boolean, we are asking whether the left or the right is True or False. If one of them is true, this will get propagated up as we recurse. 
}