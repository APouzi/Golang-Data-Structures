package stack

func isValid(s string) bool {
	var open map[rune]rune = map[rune]rune{'(': ')', '{': '}', '[': ']'}

	var stack []rune = []rune{}
	for _, v := range s {
		if _, ok := open[v]; ok {
			stack = append(stack, open[v])
		} else {
			if len(stack) == 0 {
				return false
			}
			var popped rune = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if v != popped {
				return false
			}
		}

	}

	return len(stack) <= 0
}