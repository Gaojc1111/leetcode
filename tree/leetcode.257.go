package tree

import (
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	var s string
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == node.Right { // 指针都为nil
			s += strconv.Itoa(node.Val)
			ans = append(ans, s)
			s = s[:len(s)-len(strconv.Itoa(node.Val))]
			return
		}
		s += strconv.Itoa(node.Val) + "->"
		dfs(node.Left)
		dfs(node.Right)
		s = s[:len(s)-len(strconv.Itoa(node.Val))-2]
	}
	dfs(root)
	return ans
}
