package review

// https://leetcode.cn/problems/linked-list-cycle-ii/description/
// 思路：快慢指针，快指针每次走两个节点，慢指针每次走一个节点，如果存在环，快指针必然在某一步等于慢指针
// 如果无环，则快指针会走到nil
// 设 a等于环之前的长度， k = 第一次相遇离环起点的长度， b = 整个环的长度
// 第一次相遇时，可得 a+b+k=2*(a+k) 长指针=2倍短指针 -》 a = b-k(端指针到末尾剩下的路程) 两个一起+1，就得到相遇的环起点
func detectCycle(head *ListNode) *ListNode {
	// 有环必然有两个节点及以上，因此快慢指针应该从虚拟头节点开始走
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast != slow {
		return nil
	}
	p := dummyHead
	for p != slow {
		p = p.Next
		slow = slow.Next
		if p == slow {
			return p
		}
	}
	return nil
}

//a+b+k = 2(a+k) b = a + k = cnt  a = b-k
