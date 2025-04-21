package review

// 思路：递归
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	node := root
	cnt1 := 0
	cnt2 := 0
	for node != nil {
		node = node.Left
		cnt1++
	}
	node = root
	for node != nil {
		node = node.Right
		cnt2++
	}
	if cnt1 == cnt2 && cnt1 != 0 {
		return 1<<cnt1 - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
