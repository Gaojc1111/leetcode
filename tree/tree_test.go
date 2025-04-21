package tree

import (
	"leetcode/review"
	"testing"
)

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

	//levelOrder(&TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val: 9,
	//	},
	//	Right: &TreeNode{
	//		Val: 20,
	//		Left: &TreeNode{
	//			Val: 15,
	//		},
	//		Right: &TreeNode{
	//			Val: 7,
	//		},
	//	},
	//})

	review.binaryTreePaths(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	})
}
