package tree

func postorderTraversal1(root *TreeNode) []int {
	// 递归
	if root == nil {
		return []int{}
	}

	ans := []int{}
	ans = append(ans, postorderTraversal1(root.Left)...)
	ans = append(ans, postorderTraversal1(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}

func postorderTraversal2(root *TreeNode) []int {
	// 迭代
	if root == nil {
		return []int{}
	}

	ans := []int{}
	stack := []*TreeNode{root}
	m := map[*TreeNode]bool{}

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 出栈

		if m[root] {
			ans = append(ans, root.Val)
			continue
		}

		stack = append(stack, root)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		m[root] = true
	}
	return ans
}
