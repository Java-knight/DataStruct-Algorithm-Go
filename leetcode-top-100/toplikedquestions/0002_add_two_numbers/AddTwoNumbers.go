package main

// 两数相加
// https://leetcode.cn/problems/add-two-numbers/?favorite=2cktkvj
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	num1 := 0
	num2 := 0
	sum := 0
	cur1 := l1
	cur2 := l2
	var pre *ListNode
	var node *ListNode
	for cur1 != nil || cur2 != nil {
		if cur1 != nil {
			num1 = cur1.Val
		} else {
			num1 = 0
		}
		if cur2 != nil {
			num2 = cur2.Val
		} else {
			num2 = 0
		}
		sum = num1 + num2 + carry
		pre = node
		node = &ListNode{Val: sum % 10}
		node.Next = pre
		carry = sum / 10
		if cur1 != nil {
			cur1 = cur1.Next
		} else {
			cur1 = nil
		}
		if cur2 != nil {
			cur2 = cur2.Next
		} else {
			cur2 = nil
		}
	}
	// 打补丁
	if carry == 1 {
		pre = node
		node = &ListNode{Val: carry}
		node.Next = pre
	}
	result := reverseList(node)
	return result
}

func reverseList(head *ListNode) *ListNode {
	var pre, next *ListNode
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
