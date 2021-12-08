package solution1

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
	queue, support MyQueue
}

func Constructor() MyStack {
	return MyStack{
		queue:   MyQueue{},
		support: MyQueue{},
	}
}

func (this *MyStack) Push(x int) {
	this.support.Offer(x)
	this.transferSupport2Queue()
	this.swapSupportAndQueue()
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

func (this *MyStack) transferSupport2Queue() {
	for !this.queue.Empty() {
		this.support.Offer(this.queue.Poll())
	}
}

func (this *MyStack) swapSupportAndQueue() {
	this.queue, this.support = this.support, this.queue
}
