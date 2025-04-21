package review

import (
	"leetcode/tree"
	"strconv"
)

func binaryTreePaths(root *tree.TreeNode) []string {
	ans := []string{}
	var s string
	var dfs func(node *tree.TreeNode)
	dfs = func(node *tree.TreeNode) {
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
