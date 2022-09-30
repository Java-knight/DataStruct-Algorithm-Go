package _083_remove_duplicates_from_sorted_list

// 删除排序链表中的重复元素
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pre := head
	next := head.Next
	for next != nil {
		if pre.Val == next.Val { // 出现重复
			pre.Next = next.Next
		} else { // 没有重复
			pre = next
		}
		next = pre.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
