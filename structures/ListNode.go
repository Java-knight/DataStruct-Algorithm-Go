package structures

import "fmt"

// ListNode 是链表节点
// 这个不能复制到 *_test.go 文件中。会导致 Travis 失败
type ListNode struct {
	Val  int
	Next *ListNode
}

// 链表转换为数组（切片 []int）
func List2Ints(head *ListNode) []int {
	// 链表长度限制, 链表长度超出此限制, 会 panic
	limit := 100
	index := 0

	result := []int{}
	for head != nil {
		index++
		if index > limit {
			msg := fmt.Sprintf("链表长度超出%d, 可能出现环状链表. 请检查错误, 或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// 数组转化为链表（list）
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	virHeadNode := &ListNode{} // 虚拟头节点（为了for循环好赋值）
	cur := virHeadNode
	for _, val := range nums {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return virHeadNode.Next
}

// 返回 list 第一个值为 val 的 Node
func (list *ListNode) GetNodeWith(val int) *ListNode {
	node := list
	for node != nil {
		if node.Val == val {
			break
		}
		node = node.Next
	}
	return node
}

// Ints2ListWithCycle 返回环形链表（链表尾节点.Next ——> listNode[pos]）
// 如果 pos = -1, 表示链表无环
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}

	cur := head
	for pos > 0 {
		cur = cur.Next
		pos--
	}
	tail := cur
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = cur
	return head
}
