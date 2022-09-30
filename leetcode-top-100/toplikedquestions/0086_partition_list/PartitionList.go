package _086_partition_list

// 分隔链表
// https://leetcode.cn/problems/partition-list/

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 搞两个新链表(尾插法)
	minHead := &ListNode{0, nil}
	maxHead := &ListNode{0, nil}
	min := minHead
	max := maxHead
	for head != nil {
		next := head.Next
		head.Next = nil
		if head.Val < x {
			min.Next = head
			min = min.Next
		} else {
			max.Next = head
			max = max.Next
		}
		head = next
	}
	min.Next = maxHead.Next
	return minHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
