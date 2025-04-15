package list

// 思路：双指针，间隔 n 个位置，右指针为nil时，左指针正好指向这个节点的前一个节点
// 可能删除头节点，加上虚拟头
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	i, j := dummyHead, head
	for range n {
		j = j.Next
	}

	for j != nil {
		i = i.Next
		j = j.Next
	}

	i.Next = i.Next.Next
	return dummyHead.Next
}
