package _021_merge_two_sorted_lists

// 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil || list2 == nil {
		return Ternary(list1 == nil, list2, list1)
	}
	head := Ternary(list1.Val <= list2.Val, list1, list2)
	cur1 := head.Next
	cur2 := Ternary(list1 == head, list2, list1)
	pre := head
	for cur1 != nil && cur2 != nil {
		if cur1.Val <= cur2.Val {
			pre.Next = cur1
			cur1 = cur1.Next
		} else {
			pre.Next = cur2
			cur2 = cur2.Next
		}
		pre = pre.Next
	}
	pre.Next = Ternary(cur1 == nil, cur2, cur1)
	return head
}

func Ternary(flag bool, a, b *ListNode) *ListNode {
	if flag {
		return a
	}
	return b
}

type ListNode struct {
	Val  int
	Next *ListNode
}
