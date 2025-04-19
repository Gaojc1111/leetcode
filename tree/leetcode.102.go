package tree

// 三个变量：当前正在遍历的层数，当前层数剩下的节点，下层节点数
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	ans := [][]int{}
	tmp := []int{}
	cnt, cntNext := 1, 0

	for len(queue) != 0 {
		// 加入当前节点，当前层数节点--，
		root = queue[0]
		queue = queue[1:]
		cnt--
		tmp = append(tmp, root.Val)
		if root.Left != nil {
			queue = append(queue, root.Left)
			cntNext++
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
			cntNext++
		}
		if cnt == 0 {
			cnt, cntNext = cntNext, 0
			//ans = append(ans, tmp) 不能直接append一个切片，会引用同一个切片
			t := make([]int, len(tmp))
			copy(t, tmp)
			ans = append(ans, t)
			tmp = tmp[:0]
		}
	}
	return ans
}
