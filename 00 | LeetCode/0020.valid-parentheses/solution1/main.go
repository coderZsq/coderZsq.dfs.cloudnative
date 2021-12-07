package solution1

func isValid(s string) bool {
	var stack []rune

	for _, parenthesis := range s {
		switch parenthesis {
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

			if parenthesis != stack[stackSize-1] {
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