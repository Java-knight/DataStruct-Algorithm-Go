package _110_balanced_binary_tree

import "math"

// 平衡二叉树
// https://leetcode.cn/problems/balanced-binary-tree/

var flag bool

func isBalanced(root *TreeNode) bool {
	flag = true // 表示以当前节点出发的树是否是平衡二叉树
	process(root)
	return flag
}

// 计算二叉树的最大高度
func process(node *TreeNode) int {
	if !flag || node == nil { // 已经不平衡了或 走到了叶子节点的左右孩子
		return 0
	}

	leftHeight := process(node.Left)
	rightHeight := process(node.Right)
	// 更新当前node是否平衡
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		flag = false
	}
	return Max(leftHeight, rightHeight) + 1
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
