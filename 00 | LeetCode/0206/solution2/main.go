package solution2

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
	// nil
	var newHead *ListNode
	for one != nil {
		two := one.Next
		// 1 -> nil
		// 2 -> 1 -> nil
		one.Next = newHead
		// newHead 1
		// newHead 2
		newHead = one

		one = two
	}
	return newHead
}

func done(head *ListNode) bool {
	return head == nil || head.Next == nil
}