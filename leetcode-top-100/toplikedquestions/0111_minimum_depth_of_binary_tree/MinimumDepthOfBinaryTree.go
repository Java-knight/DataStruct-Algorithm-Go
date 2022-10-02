package _111_minimum_depth_of_binary_tree

// 二叉树的最小深度
// https://leetcode.cn/problems/minimum-depth-of-binary-tree/

// BFS 实现(时间复杂度O(N), 空间复杂度O(N))
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 1
	var queue []*TreeNode
	size := 0
	queue = append(queue, root)
	for len(queue) != 0 {
		size = len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return result
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result++
	}
	return result
}

//// 暴力递归
//func minDepth(root *TreeNode) int {
//	if root == nil {  // base case1 走到空, 叶子节点的左右孩子(其中一种)
//		return 0
//	}
//	// base case2 当前节点不是null, 如果左子树或右子树最小高度为0, 表示当前node是叶子节点
//	left := minDepth(root.Left)
//	right := minDepth(root.Right)
//	if left == 0 || right == 0 {
//		return left + right + 1
//	}
//	// base case3 当前节点是非叶子节点, 同时有左右孩子
//	return Min(left, right) + 1
//}
//
//func Min(a, b int) int {
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
