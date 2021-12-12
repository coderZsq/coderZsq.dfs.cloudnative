package solution1

import . "_/structures/binary-tree-node"

func isValidBST(root *TreeNode) bool {
	nums := inorder(root)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}

func inorder(root *TreeNode) (result []int) {
	if root == nil {
		return []int{}
	}
	result = inorder(root.Left)
	result = append(result, root.Val)
	result = append(result, inorder(root.Right)...)
	return
}
