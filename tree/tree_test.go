package tree

import "testing"

func TestTree(t *testing.T) {
	//preorderTraversal2(&TreeNode{
	//	Val: 1,
	//	Right: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 3,
	//		},
	//	},
	//})

	levelOrder(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	})
}
