package list

// 思路：快慢指针，快指针每次走两个节点，慢指针每次走一个节点，如果存在环，快指针必然在某一步等于慢指针
func hasCycle(head *ListNode) bool {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
