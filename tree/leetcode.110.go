package tree

import "math"

// 思路：每个节点都要判断，递归。 先处理子树节点，再处理自身
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	if math.Abs(float64(maxDepth1(root.Left)-maxDepth1(root.Right))) <= 1 {
		return true
	}
	return false
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth1(root.Left), maxDepth1(root.Right)) + 1
}
