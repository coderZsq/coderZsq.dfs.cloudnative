package solution2

import . "_/structures/single-linkedlist"

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	zero := dummyHead
	for zero.Next != nil && zero.Next.Next != nil {
		one := zero.Next
		two := zero.Next.Next
		// swap one
		rebind(zero, one, two)
		zero = one
	}
	return dummyHead.Next
}

// 0 -> 1 -> 2 -> nil
func rebind(zero *ListNode, one *ListNode, two *ListNode) {
	// 0 -> 2
	zero.Next = two
	// 1 -> nil
	one.Next = two.Next
	// 0 -> 2 -> 1 -> nil
	two.Next = one
}
