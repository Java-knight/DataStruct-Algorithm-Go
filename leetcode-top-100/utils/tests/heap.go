package tests

// 堆的实现
// 可以参考 sdk/container/heap/example_pq_test.go的实现
type ListNode struct {
	Val  int
	Next *ListNode
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
