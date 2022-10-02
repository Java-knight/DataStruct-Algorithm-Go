package _116_populating_next_right_pointers_in_each_node

// 填充每个节点的下一个右侧节点指针
// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/

// 层序遍历(自己实现一个queue), 空间复杂度O(1); 放好这个Node有next, 自己实现的queue就两个指针
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := &queue{nil, nil, 0}
	queue.offer(root)
	var pre *Node
	var size int
	for !queue.isEmpty() {
		pre = nil
		size = queue.size
		for i := 0; i < size; i++ {
			node := queue.poll()
			if pre == nil { // 该level第一次访问
				pre = node
			} else { // 该level 非第一次访问
				pre.Next = node
				pre = node
			}
			if node.Left != nil {
				queue.offer(node.Left)
			}
			if node.Right != nil {
				queue.offer(node.Right)
			}
		}
	}
	return root
}

type queue struct {
	head *Node
	tail *Node
	size int
}

func (q *queue) offer(val *Node) {
	q.size++
	if q.head == nil {
		q.head = val
		q.tail = val
	} else {
		q.tail.Next = val
		q.tail = val
	}
}

func (q *queue) poll() *Node {
	q.size--
	node := q.head
	q.head = q.head.Next
	node.Next = nil
	return node
}

func (q *queue) isEmpty() bool {
	return q.size == 0
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
