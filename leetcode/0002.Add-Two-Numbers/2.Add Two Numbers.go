package leetcode

import "DataStruct-Algorithm-Go/structures"

// ListNode define
type ListNode = structures.ListNode

// https://leetcode.cn/problems/add-two-numbers/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	// 一个虚拟头节点（最后结果链表 result.Next）
	result := &ListNode{Val: 0}
	cur := result // 在链表上跑
	carry := 0    // 进位

	val1, val2 := 0, 0
	for l1 != nil || l2 != nil || carry != 0 {
		// 注意：这里不能使用三目运算
		if l1 == nil {
			val1 = 0
		} else {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			val2 = 0
		} else {
			val2 = l2.Val
			l2 = l2.Next
		}
		sum := val1 + val2 + carry
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
	}
	return result.Next
}
