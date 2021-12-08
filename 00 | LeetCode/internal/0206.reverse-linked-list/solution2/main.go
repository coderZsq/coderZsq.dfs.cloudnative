package solution2

import . "_/structures/single-linkedlist"

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		return reverse(head)
	}
}

func reverse(node *ListNode) *ListNode {
	var newHead *ListNode
	for node != nil {
		next := node.Next
		rebind(&node, &newHead)
		node = next
	}
	return newHead
}

func rebind(node **ListNode, newHead **ListNode) {
	// 1 -> nil
	// 2 -> 1 -> nil
	(*node).Next = *newHead
	// newHead 1
	// newHead 2
	*newHead = *node
}