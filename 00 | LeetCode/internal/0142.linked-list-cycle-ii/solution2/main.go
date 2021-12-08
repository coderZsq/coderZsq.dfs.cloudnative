package solution2

import . "_/structures/single-linkedlist"

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return detect(head, slow)
		}
	}
	return nil
}

func detect(head *ListNode, slow *ListNode) *ListNode {
	node := head
	for node != slow {
		node = node.Next
		slow = slow.Next
	}
	return node
}
