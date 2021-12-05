package solution2

import (
	. "_/definition/single-linkedlist"
)

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	zero := dummyHead
	for !done(zero.Next) {
		one := zero.Next
		// swap one
		swap(zero, one, one.Next)
		zero = one
	}
	return dummyHead.Next
}

func done(head *ListNode) bool {
	return head == nil || head.Next == nil
}

// 0 -> 1 -> 2 -> nil
func swap(zero *ListNode, one *ListNode, two *ListNode) {
	// 0 -> 2
	zero.Next = two
	// 1 -> nil
	one.Next = two.Next
	// 0 -> 2 -> 1 -> nil
	two.Next = one
}
