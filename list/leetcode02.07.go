package list

// 思路：双指针，把两个链表拼接在一起，总长度一样，如果有相交，那么指针一定相同。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB

	// 退出有两种情况，1. a和b相交 2. a和b走完，都为nil
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}
