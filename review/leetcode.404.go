package tree

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			return sumOfLeftLeaves(root.Right) + root.Left.Val
		} else {
			return sumOfLeftLeaves(root.Right) + sumOfLeftLeaves(root.Left)
		}
	}
	return sumOfLeftLeaves(root.Right)
}
