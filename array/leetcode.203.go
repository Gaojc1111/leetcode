package array

type ListNode struct {
	Val  int
	Next *ListNode
}

// 因为可能会删除头节点，所以需要设置一个虚拟头节点。
func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}

	i, j := dummyHead, head

	for j != nil {
		if j.Val == val {
			i.Next = j.Next
			j = j.Next
			continue
		}
		i = j
		j = j.Next
	}

	return dummyHead.Next
}
