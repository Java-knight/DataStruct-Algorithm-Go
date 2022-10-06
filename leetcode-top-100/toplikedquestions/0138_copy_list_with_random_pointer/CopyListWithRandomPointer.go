package _138_copy_list_with_random_pointer

// 复制带随机指针的链表
// https://leetcode.cn/problems/copy-list-with-random-pointer/

// 思路: 将老链表每个节点后面都赋值一份, 设置好random指针, 进行拆分
// https://www.processon.com/view/link/633e81ba6376891c6b42067e
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var cur, next, copyCur *Node
	cur = head
	// (1) 复制链表 1—>2—>3—>4    1—>1'—>2—>2'—>3—>3'—>4—>4'
	for cur != nil {
		next = cur.Next
		cur.Next = &Node{cur.Val, nil, nil}
		cur.Next.Next = next
		cur = next
	}

	// (2) 给复制的节点rand指针赋值
	cur = head
	for cur != nil {
		next = cur.Next.Next
		copyCur = cur.Next
		if cur.Random != nil {
			copyCur.Random = cur.Random.Next
		}
		cur = next
	}

	// 拆链表
	result := head.Next
	cur = head
	for cur != nil {
		next = cur.Next.Next
		copyCur = cur.Next
		if next != nil {
			copyCur.Next = next.Next
		} else {
			copyCur.Next = nil
		}
		cur.Next = next
		cur = next
	}
	return result
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
