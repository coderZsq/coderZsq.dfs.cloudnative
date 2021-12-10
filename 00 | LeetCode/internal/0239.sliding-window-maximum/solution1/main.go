package main

import (
	"container/heap"
	"sort"
)

type MyHeap struct {
	sort.IntSlice
	k int
}

func (this *MyHeap) Less(i, j int) bool {
	return this.IntSlice[i] > this.IntSlice[j]
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

func maxSlidingWindow(nums []int, k int) []int {
	max := &MyHeap{IntSlice: []int{}, k: k}
	var result []int
	preKProcess(&result, max, nums)
	for i := k; i < len(nums); i++ {
		index := findIdxIn(max.IntSlice, nums[i-k])
		heap.Remove(max, index)
		heap.Push(max, nums[i])
		result = append(result, max.Peek())
	}
	return result
}

func preKProcess(result *[]int, max *MyHeap, nums []int) {
	*result = append(*result, findTopK(max, nums))
}

func findTopK(max *MyHeap, nums []int) (peek int) {
	for i := 0; i < max.k; i++ {
		heap.Push(max, nums[i])
	}
	peek = max.Peek()
	return
}

func findIdxIn(arr []int, poll int) (index int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == poll {
			index = i
		}
	}
	return
}
