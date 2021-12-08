package solution2

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

func (this *MyStack) Push(x int) {
	queueSize := len(this.queue)
	this.queueOfferTail(x)
	this.pollQueueHeadOfferToQueueTail(queueSize)
}

func (this *MyStack) Pop() int {
	queueHead := this.queue[0]
	this.queuePollHead()
	return queueHead
}

func (this *MyStack) Top() int {
	return this.queue[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

func (this *MyStack)pollQueueHeadOfferToQueueTail(size int)  {
	for i := 0; i < size; i++ {
		queueHead := this.queue[0]
		this.queuePollHead()
		this.queueOfferTail(queueHead)
	}
}

func (this *MyStack)queuePollHead()  {
	this.queue = this.queue[1:]
}

func (this *MyStack) queueOfferTail(x int)  {
	this.queue = append(this.queue, x)
}