package _019_remove_nth_node_from_end_of_list

// 删除莲表的倒数第N个节点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var pre *ListNode
	cur := head
	flag := false
	for cur != nil {
		n--
		if n <= 0 {
			if n == 0 { // 存在倒数第N个节点
				flag = true
			} else if n == -1 { // 找到了倒数第N+1个节点
				pre = head
			} else {
				pre = pre.Next
			}
		}
		cur = cur.Next
	}
	if !flag { // base case1 链表上不存在倒数第N个节点
		return head
	}
	if pre == nil { // base case2 需要删除的第N个几点就是头节点
		return head.Next
	}
	pre.Next = pre.Next.Next // base case3 正常情况
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//// 可以省略一个变量（面试吹牛逼的，没啥用）
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	var pre *ListNode
//	cur := head
//	for cur != nil {
//		n--
//		if n == -1 {  // 找到了倒数第N+1个节点
//			pre = head
//		}
//		if n < -1 {
//			pre = pre.Next
//		}
//
//		cur = cur.Next
//	}
//	if n > 0 {  // base case1 链表上不存在倒数第N个节点
//		return head
//	}
//	if pre == nil {  // base case2 需要删除的第N个几点就是头节点
//		return head.Next
//	}
//	pre.Next = pre.Next.Next  // base case3 正常情况
//	return head
//}
