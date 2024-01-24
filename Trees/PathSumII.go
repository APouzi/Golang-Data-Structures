package bst

// Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths where the sum of the node values in the path equals targetSum. Each path should be returned as a list of the node values, not node references.
// A root-to-leaf path is a path starting from the root and ending at any leaf node. A leaf is a node with no children.

//Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// Output: [[5,4,11,2],[5,8,4,5]]
// Explanation: There are two paths whose sum equals targetSum:
// 5 + 4 + 11 + 2 = 22
// 5 + 8 + 4 + 5 = 22

//Example 2:
// Input: root = [1,2,3], targetSum = 5
// Output: []

//Example 3:
// Input: root = [1,2], targetSum = 0
// Output: []

func PathSumII(root *Node, targetSum int) [][]int {
    if root == nil{
        return [][]int{}
    }
    var ans *[][]int = &[][]int{}
    dfsPathSum(root, targetSum,  ans, []int{})
    return *ans
}

func dfsPathSum(root *Node, targetSum int, ans *[][]int, temp []int){
    if root == nil{
        return 
    }
    temp = append(temp, root.val)
    if root != nil && (root.left == nil && root.right == nil) && targetSum - root.val == 0{
        // temp = append(temp, root.val)
        *ans = append(*ans, append([]int{},temp...)) //What we need to understand is that we can't just "*ans = append(*ans, temp)", this is because temp is a refrence to the slice. So if that slice get's change as its recursed back to a previous state and then appeneded to, that appeneded slice will be changed. Remember, it's NOT just a value.  
        return 
    }
    
    
    
    dfsPathSum(root.left, targetSum - root.val, ans,temp)
    dfsPathSum(root.right, targetSum - root.val, ans,temp) 
}







func pathSumIIV2(root *Node, targetSum int) [][]int {
    if root == nil{
        return [][]int{}
    }
    var ans [][]int = [][]int{}
    ret := dfsPathSumV2(root, targetSum,  ans, []int{})
    return ret
}

func dfsPathSumV2(root *Node, targetSum int, ans [][]int, temp []int)[][]int{
    if root == nil{
        return ans
    }
    temp = append(temp, root.val)
    if root != nil && (root.left == nil && root.right == nil) && targetSum - root.val == 0{
        // temp = append(temp, root.Val)
        ans = append(ans, append([]int{},temp...))
        return ans
    }
    
    left := dfsPathSumV2(root.left, targetSum - root.val, ans,temp)
    right := dfsPathSumV2(root.right, targetSum - root.val, ans,temp)
    ans = append(ans, left...)
    ans = append(ans, right...)
    

    return ans
}