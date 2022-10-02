package _109_convert_sorted_list_to_binary_search_tree

// 有序链表转换搜索二叉树
// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/

// 思路: 使用二分法, 将list转换为 array
func sortedListToBST(head *ListNode) *TreeNode {
	// list转array
	cur := head
	arr := make([]int, 0)
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, left, right int) *TreeNode {
	if left > right { // base case1 叶子节点的左右孩子
		return nil
	}

	// base case2 叶子/非叶子节点的创建
	mid := left + ((right - left) >> 1)
	node := &TreeNode{arr[mid], nil, nil}
	node.Left = process(arr, left, mid-1)
	node.Right = process(arr, mid+1, right)
	return node
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
