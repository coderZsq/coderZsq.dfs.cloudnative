package solution1

import (
	. "_/definition/single-linkedlist"
)

func reverseList(head *ListNode) *ListNode {
	if done(head) {
		return head
	} else {
		return reverse(head)
	}
}

// 1 -> 2 -> 3 -> 4 -> 5
func reverse(one *ListNode) *ListNode {
	// newHead 5 -> 4 -> 3 -> 2
	newHead := reverseList(one.Next)
	// one               1 -> 2
	adjust(one)
	// newHead 5 -> 4 -> 3 -> 2 -> 1 -> nil
	return newHead
}

// 1 -> 2
func adjust(one *ListNode) {
	two := one.Next
	// 2 -> 1
	two.Next = one
	// 1 -> nil
	one.Next = nil
}

func done(head *ListNode) bool {
	return head == nil || head.Next == nil
}