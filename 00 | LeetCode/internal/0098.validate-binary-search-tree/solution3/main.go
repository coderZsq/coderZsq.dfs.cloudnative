package solution3

import (
	. "_/structures/binary-tree-node"
	"math"
)

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min {
		return false
	}
	if root.Val >= max {
		return false
	}
	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}
