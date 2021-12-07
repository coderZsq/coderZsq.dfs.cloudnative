package solution2

import . "_/structures/single-linkedlist"

func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			node := head

			for node != slow {
				slow = slow.Next
				node = node.Next
			}

			return node
		}
	}
	return nil
}
