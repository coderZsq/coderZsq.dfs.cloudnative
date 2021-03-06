package solution1

import . "_/structures/single-linkedlist"

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}

	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(node *ListNode, k *ListNode) *ListNode {
	newHead := k

	for node != k {
		next := node.Next
		rebind(&node, &newHead)
		node = next
	}

	return newHead
}

func rebind(node **ListNode, newHead **ListNode)  {
	(*node).Next = *newHead
	*newHead = *node
}