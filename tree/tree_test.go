package tree

import "testing"

func TestTree(t *testing.T) {
	preorderTraversal2(&TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	})
}
