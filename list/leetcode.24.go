package list

// 因为会修改头节点，创建虚拟头节点
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}

	// 核心是a和b互换
	pre := dummyHead
	a := head

	// a,b 不能为nil
	for a != nil && a.Next != nil {
		b := a.Next
		c := b.Next

		pre.Next = b
		a.Next = c
		b.Next = a

		// 迭代
		// 第一次迭代得到 pre-2-1-3-4
		pre = a
		a = pre.Next
	}

	return dummyHead.Next
}
