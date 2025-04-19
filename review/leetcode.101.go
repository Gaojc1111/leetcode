package review

// 思路：不要去思考一个节点一个节点对比，而是一棵子树一棵子树对比
// 每次递归两颗子树
func isSymmetric(root *TreeNode) bool {
	return check1(root.Left, root.Right)
}

func check1(l, r *TreeNode) bool {
	// 1. 递归
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil {
		return false
	}

	// l, r自身相等，左右子树分别对称
	if l.Val != r.Val {
		return false
	}

	if check1(l.Left, r.Right) && check1(l.Right, r.Left) {
		return true
	}
	return false
}

func check2(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	// 2.迭代
	// 两颗子树分别走一次遍历，按照顺序比较
	stack1 := []*TreeNode{l}
	stack2 := []*TreeNode{r}

	arr1 := []int{}
	arr2 := []int{}

	for len(stack1) != 0 {
		root := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]

		arr1 = append(arr1, root.Val)

		if root.Left == nil && root.Right == nil {
			continue
		}
		if root.Left != nil && root.Right != nil {
			stack1 = append(stack1, root.Left)
			stack1 = append(stack1, root.Right)
		} else if root.Left == nil {
			stack1 = append(stack1, &TreeNode{Val: -111})
			stack1 = append(stack1, root.Right)
		} else if root.Right == nil {
			stack1 = append(stack1, root.Left)
			stack1 = append(stack1, &TreeNode{Val: -111})
		}
	}

	for len(stack2) != 0 {
		root := stack2[len(stack2)-1]
		stack2 = stack2[:len(stack2)-1]

		arr2 = append(arr2, root.Val)

		if root.Left == nil && root.Right == nil {
			continue
		}
		if root.Right != nil && root.Left != nil {
			stack2 = append(stack2, root.Right)
			stack2 = append(stack2, root.Left)
		} else if root.Right == nil {
			stack2 = append(stack2, &TreeNode{Val: -111})
			stack2 = append(stack2, root.Left)
		} else if root.Left == nil {
			stack2 = append(stack2, root.Right)
			stack2 = append(stack2, &TreeNode{Val: -111})
		}
	}

	if len(arr1) != len(arr2) {
		return false
	}
	for k, _ := range arr1 {
		if arr1[k] != arr2[k] {
			return false
		}
	}
	return true
}
