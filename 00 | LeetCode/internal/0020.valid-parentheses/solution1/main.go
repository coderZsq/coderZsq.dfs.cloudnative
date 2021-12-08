package solution1

func isValid(s string) bool {
	var stack []rune

	for _, parentheses := range s {
		switch parentheses {
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case ')', '}', ']':
			stackSize := len(stack)

			if empty(stack) {
				return false
			}

			if parentheses != stack[stackSize-1] {
				return false
			}

			stack = stack[:stackSize-1]
		}
	}
	return empty(stack)
}

func empty(stack []rune) bool {
	return len(stack) == 0
}