package _082_remove_duplicates_from_sorted_list_ii

// 删除链表中的重复元素II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := &ListNode{0, head} // 虚拟头节点
	cur := result
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val { // 出现了重复元素
			val := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == val {
				cur.Next = cur.Next.Next
			}
		} else { // 没有重复元素
			cur = cur.Next
		}
	}
	return result.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
