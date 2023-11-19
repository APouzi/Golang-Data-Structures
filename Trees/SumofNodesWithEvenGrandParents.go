package bst

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

func sumEvenGrandparent(root *Node) int {
	if root == nil || root.right == nil && root.left == nil{
		return 0
	}
	stackGrandPa := []int{}
	ans := 0
	return DFSGrandParentSum(root, stackGrandPa, &ans)
}

func DFSGrandParentSum(root *Node, stack []int,ans *int ) int{
	if root == nil{
		return 0
	}
	stack = append(stack,root.val)
	DFSGrandParentSum(root.left,stack,ans)
	DFSGrandParentSum(root.right,stack,ans)

	if len(stack) > 2 && isEven(stack[len(stack)-3]){
		*ans += root.val
	}

	return *ans

}

func isEven(num int)bool{
	return num % 2 == 0
}