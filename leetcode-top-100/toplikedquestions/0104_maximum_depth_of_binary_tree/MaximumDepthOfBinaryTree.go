package _104_maximum_depth_of_binary_tree

// 二叉树的最大深度
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

func maxDepth(root *TreeNode) int {
	if root == nil { // 叶子节点
		return 0
	}
	// 可以不加(判断左右孩子都是叶子节点)
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// 当前节点不是叶子节点
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
