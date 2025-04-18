package tree

func inorderTraversal1(root *TreeNode) []int {
	// 1.递归
	if root == nil {
		return []int{}
	}

	ans := []int{}
	ans = append(ans, inorderTraversal1(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversal1(root.Right)...)

	return ans
}

func inorderTraversal2(root *TreeNode) []int {
	// 2.迭代
	if root == nil {
		return []int{}
	}

	stack := []*TreeNode{root}
	ans := []int{}
	m := map[*TreeNode]bool{}

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if m[root] {
			ans = append(ans, root.Val)
			continue
		}

		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		stack = append(stack, root)
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		m[root] = true
	}
	return ans
}
