package solution2

type MyQueue struct {
	container []int
}

func (this *MyQueue) Offer(x int) {
	this.container = append(this.container, x)
}

func (this *MyQueue) Poll() int {
	peek := this.Peek()
	this.container = this.container[1:]
	return peek
}

func (this *MyQueue) Peek() int {
	peek := this.container[0]
	return peek
}

func (this *MyQueue) Empty() bool {
	return len(this.container) <= 0
}

func (this *MyQueue) Size() int {
	return len(this.container)
}

type MyStack struct {
	queue MyQueue
}

func Constructor() MyStack {
	return MyStack{
		queue: MyQueue{},
	}
}

func (this *MyStack) Push(x int) {
	size := this.queue.Size()
	this.queue.Offer(x)
	this.transferRange(size)
}

func (this *MyStack) Pop() int {
	return this.queue.Poll()
}

func (this *MyStack) Top() int {
	return this.queue.Peek()
}

func (this *MyStack) Empty() bool {
	return this.queue.Empty()
}

func (this *MyStack) transferRange(size int) {
	for i := 0; i < size; i++ {
		this.queue.Offer(this.queue.Poll())
	}
}
