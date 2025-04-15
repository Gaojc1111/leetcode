package list

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 1. 迭代，创建虚拟头节点，一个一个反转
	dummyHead := &ListNode{Next: head}
	pre, cur := dummyHead, head
	for cur != nil {
		ne := cur.Next

		cur.Next = pre

		// 迭代
		pre = cur
		cur = ne
	}
	dummyHead.Next.Next = nil // 别忘了让 1号节点 指向nil
	return pre

	//// 2. 递归
	//// 思路: 递归结束条件 节点为nil
	//// 返回值：翻转好的链表头节点，5-4-3，返回5
	//// 递归处理：将当前节点比如2，加入到返回到头节点链表中，通过2->3这个关系轻松拿到尾节点
	//
	//root := reverseList(head.Next)
	//if root == nil {
	//	return head
	//}
	//head.Next.Next = head // 加入尾节点
	//head.Next = nil       // 尾节点指向nil
	//return root
}
