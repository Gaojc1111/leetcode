package review

/*
https://leetcode.cn/problems/reverse-linked-list-ii/description/
总体思路：双指针，类似反转链表,计算迭代次数
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Next: head}

	// 定位到Left上一个
	a := dummyHead
	for i := 1; i <= left-1; i++ {
		a = a.Next
	}

	// 反转链表
	pre := a.Next // pre == left, 不会为nil
	cur := pre.Next

	for i := 1; i <= right-left; i++ {
		ne := cur.Next

		cur.Next = pre

		pre, cur = cur, ne
	}

	a.Next.Next = cur // 2->5
	a.Next = pre      // 1->4
	return dummyHead.Next
}
