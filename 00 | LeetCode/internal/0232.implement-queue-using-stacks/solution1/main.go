package solution1

type MyStack struct {
	container []int
}

func (this *MyStack) Push(x int) {
	this.container = append(this.container, x)
}

func (this *MyStack) Pop() int {
	top := this.Top()
	this.container = this.container[:len(this.container)-1]
	return top
}

func (this *MyStack) Top() int {
	top := this.container[len(this.container)-1]
	return top
}

func (this *MyStack) Empty() bool {
	return len(this.container) <= 0
}

type MyQueue struct {
	input, output MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		input:  MyStack{},
		output: MyStack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.input.Push(x)
}

func (this *MyQueue) Pop() int {
	this.check()
	return this.output.Pop()
}

func (this *MyQueue) Peek() int {
	this.check()
	return this.output.Top()
}

func (this *MyQueue) Empty() bool {
	return this.input.Empty() && this.output.Empty()
}

func (this *MyQueue) check() {
	if !this.input.Empty() && this.output.Empty() {
		this.transfer()
	}
}

func (this *MyQueue) transfer() {
	for !this.input.Empty() {
		this.output.Push(this.input.Pop())
	}
}
