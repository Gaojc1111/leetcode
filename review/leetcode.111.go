package review

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if l == 0 && r == 0 {
		return 1
	}
	if l == 0 {
		return r + 1
	}
	if r == 0 {
		return l + 1
	}
	return min(l, r) + 1
}
