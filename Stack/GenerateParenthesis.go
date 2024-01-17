package stack

//Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

//Example 1:
//Input: n = 3
//Output: ["((()))","(()())","(())()","()(())","()()()"]

//Example 2:
//Input: n = 1
//Output: ["()"]

func generateParenthesis(n int) []string {
	var ans []string = []string{}
	var open, closed int = 0, 0
	var stack []byte = make([]byte, 0)
	
	BackTracking(stack, &ans, open, closed, n)
    return ans
}

 func BackTracking(stack []byte, ans *[]string, open, closed, n int){
	if open == n && closed == n{
		*ans = append(*ans, string(stack))
		return
	}

	if open < n{
		stack = append(stack, '(')
		BackTracking(stack,ans, open+1, closed,n)
		stack = stack[:len(stack)-1]
	}

	if closed < open {
		stack = append(stack, ')')
		BackTracking(stack,ans, open, closed+1,n)
		stack = stack[:len(stack)-1]
	}
}




func generateParenthesisV2(n int) []string {
	var ans []string = []string{}
	var open, closed int = 0, 0
	
	BackTrackingV2("", &ans, open, closed, n)
    return ans
}

 func BackTrackingV2(stack string, ans *[]string, open, closed, n int){
	if len(stack) == n *2{
		*ans = append(*ans, stack)
		return
	}

	if open < n{
		BackTrackingV2(stack + "(",ans, open+1, closed,n)
	}

	if closed < open {
		BackTrackingV2(stack + ")",ans, open, closed+1,n)
	}
    return
}