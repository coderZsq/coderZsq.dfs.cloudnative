package solution1

import . "_/structures/single-linkedlist"

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		return reverse(head)
	}
}

// 1 -> 2 -> 3 -> 4 -> 5
func reverse(head *ListNode) *ListNode {
	// newHead 5 -> 4 -> 3 -> 2
	newHead := reverseList(head.Next)
	// head              1 -> 2
	rebind(head)
	// newHead 5 -> 4 -> 3 -> 2 -> 1 -> nil
	return newHead
}

// 1 -> 2
func rebind(one *ListNode) {
	two := one.Next
	// 2 -> 1
	two.Next = one
	// 1 -> nil
	one.Next = nil
}