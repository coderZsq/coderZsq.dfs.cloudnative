package solution1

type MyStack struct {
	container []rune
}

func (this *MyStack) Push(x rune) {
	this.container = append(this.container, x)
}

func (this *MyStack) Pop() rune {
	top := this.Top()
	this.container = this.container[:len(this.container)-1]
	return top
}

func (this *MyStack) Top() rune {
	top := this.container[len(this.container)-1]
	return top
}

func (this *MyStack) Empty() bool {
	return len(this.container) <= 0
}

func isValid(s string) bool {
	stack := &MyStack{[]rune{}}
	for _, parentheses := range s {
		switch parentheses {
		case '(':
			stack.Push(')')
		case '{':
			stack.Push('}')
		case '[':
			stack.Push(']')
		case ')', '}', ']':
			if stack.Empty() {
				return false
			}
			if parentheses != stack.Top() {
				return false
			}
			stack.Pop()
		}
	}
	return stack.Empty()
}