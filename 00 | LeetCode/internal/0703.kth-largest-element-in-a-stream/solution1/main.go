package main

import (
	"container/heap"
	"sort"
)

type MyHeap struct {
	sort.IntSlice
	k int
}

func (this *MyHeap) Push(v interface{}) {
	this.IntSlice = append(this.IntSlice, v.(int))
}

func (this *MyHeap) Pop() interface{} {
	pop := this.IntSlice[len(this.IntSlice)-1]
	this.IntSlice = this.IntSlice[:len(this.IntSlice)-1]
	return pop
}

func (this *MyHeap) Peek() int {
	return this.IntSlice[0]
}

func (this *MyHeap) Len() int {
	return len(this.IntSlice)
}

type KthLargest struct {
	queue *MyHeap
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{
		queue: &MyHeap{IntSlice: []int{}, k: k},
	}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	if this.queue.Len() < this.queue.k {
		heap.Push(this.queue, val)
	} else if this.queue.Peek() < val {
		heap.Push(this.queue, val)
		heap.Pop(this.queue)
	}
	return this.queue.Peek()
}
