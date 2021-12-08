package solution1

type MyStack struct {
	queue1, queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		queue1: []int{},
		queue2: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.queue2OfferTail(x)
	this.pollQueue1HeadOfferToQueue2Tail()
	this.swapQueues()
}

func (this *MyStack) Pop() int {
	top := this.queue1[0]
	this.queue1Poll()
	return top
}

func (this *MyStack) Top() int {
	return this.queue1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}

func (this *MyStack) queue2OfferTail(x int) {
	this.queue2 = append(this.queue2, x)
}

func (this *MyStack) pollQueue1HeadOfferToQueue2Tail() {
	for this.queue1NotEmpty() {
		this.queue2OfferTail(this.queue1[0])
		this.queue1Poll()
	}
}

func (this *MyStack) queue1NotEmpty() bool {
	return len(this.queue1) > 0
}

func (this *MyStack) queue1Poll() {
	this.queue1 = this.queue1[1:]
}

func (this *MyStack) swapQueues() {
	this.queue1, this.queue2 = this.queue2, this.queue1
}
