package main

type MyDeque struct {
	container []int
}

func (this *MyDeque) OfferLast(x int) {
	this.container = append(this.container, x)
}

func (this *MyDeque) PollFirst() int {
	peek := this.PeekFirst()
	this.container = this.container[1:]
	return peek
}

func (this *MyDeque) PollLast() int {
	peek := this.PeekFirst()
	this.container = this.container[:this.Size()-1]
	return peek
}

func (this *MyDeque) PeekFirst() int {
	peek := this.container[0]
	return peek
}

func (this *MyDeque) PeekLast() int {
	peek := this.container[this.Size()-1]
	return peek
}

func (this *MyDeque) Empty() bool {
	return this.Size() <= 0
}

func (this *MyDeque) Size() int {
	return len(this.container)
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	window := &MyDeque{
		container: []int{},
	}
	var result []int
	for i, x := range nums {
		if i >= k && window.PeekFirst() <= i-k {
			window.PollFirst()
		}
		for window.Size() > 0 && nums[window.PeekLast()] < x {
			window.PollLast()
		}
		window.OfferLast(i)
		if i >= k-1 {
			result = append(result, nums[window.PeekFirst()])
		}
	}
	return result
}
