package _023_merge_k_sorted_lists

import (
	"container/heap"
)

// 合并K个升序链表
// https://leetcode.cn/problems/merge-k-sorted-lists/

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	// (1) 把每个链表的头节点加入 minHeap 中
	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	// (2) 判断 minHeap 是否为空
	if minHeap.IsEmpty() {
		return nil
	}

	// (3) 弹出加入resultList, 将弹出的节点.next加入minHeap
	head := heap.Pop(minHeap).(*ListNode)
	pre := head
	var cur *ListNode
	if pre.Next != nil {
		heap.Push(minHeap, pre.Next)
	}
	for !minHeap.IsEmpty() {
		cur = heap.Pop(minHeap).(*ListNode)
		pre.Next = cur
		pre = cur
		if cur.Next != nil {
			heap.Push(minHeap, cur.Next)
		}
	}
	return head
}

type MinHeap []*ListNode

func (h MinHeap) Len() int {
	return len(h)
}

// 小根堆的实现
func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	val := x.(*ListNode)
	*h = append(*h, val)
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	size := len(old)
	val := old[size-1]
	old[size-1] = nil // 防止发生内存泄露
	*h = old[0 : size-1]
	return val
}

// IsEmpty: 判断是否为空, 自己扩充的函数
func (h *MinHeap) IsEmpty() bool {
	return len(*h) == 0
}

var _ heap.Interface = (*MinHeap)(nil)

type ListNode struct {
	Val  int
	Next *ListNode
}
