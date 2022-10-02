package _108_convert_sorted_array_to_binary_search_tree

// 将有序数组转换为二叉搜索树
// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/

// 思路: 二分查找的可以尽可能的让它平衡
func sortedArrayToBST(nums []int) *TreeNode {
	return process(nums, 0, len(nums)-1)
}

func process(nums []int, left, right int) *TreeNode {
	if left > right { // base case1 越界(叶子节点的左右节点)
		return nil
	}
	if left == right { // base case2 叶子节点
		return &TreeNode{nums[left], nil, nil}
	}

	// base case3 nums[mid]位置创建一个node, 继续左右连边一起创建
	mid := left + ((right - left) >> 1)
	node := &TreeNode{nums[mid], nil, nil}
	node.Left = process(nums, left, mid-1)
	node.Right = process(nums, mid+1, right)
	return node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
