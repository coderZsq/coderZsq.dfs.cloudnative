package main

import (
	. "_/structures/binary-tree-node"
)

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode = nil
	return inorder(root, &prev)
}

func inorder(root *TreeNode, prev **TreeNode) bool {
	if root == nil {
		return true
	}
	if !inorder(root.Left, prev) {
		return false
	}
	if (*prev) != nil && (*prev).Val >= root.Val {
		return false
	}
	*prev = root
	return inorder(root.Right, prev)
}
