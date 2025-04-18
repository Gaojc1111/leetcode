package tree

func preorderTraversal1(root *TreeNode) []int {
	// 1.递归
	if root == nil {
		return []int{}
	}
	ans := []int{}
	ans = append(ans, root.Val)
	ans = append(ans, preorderTraversal1(root.Left)...)
	ans = append(ans, preorderTraversal1(root.Right)...)
	return ans
}

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// 2.迭代: 维护一个节点队列
	stack := []*TreeNode{root}
	ans := []int{}

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 出栈

		ans = append(ans, root.Val)

		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}

	return ans
}
