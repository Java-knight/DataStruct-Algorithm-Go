package _687_longest_univalue_path

import (
	"math"
)

// 最长同值路径
// https://leetcode.cn/problems/longest-univalue-path/

// 二叉树的递归套路思路: x无关、x相关
// x无关:
// (1) 左子树的最大路径;  此时, max = 左子树的最大路径
// (2) 右子树的最大路径;  此时, max = 右子树的最大路径
// x有关:
// (1) 左子树和右子树都没有最大路径（x是叶子节点）, max = 1
// (2) left.val == node.val && node.val != right.val 此时, max = 左子树的以left开始的最大路径 + 1
// (3) left.val != node.val && node.val == right.val 此时, max = 右子树的以right开始的最大路径 + 1
// (4) left.val == node.val && node.val == right.val 此时, max = 右子树的以right开始的最大路径 + 左子树的以left开始的最大路径 + 1
// 结论: 要获取x节点的最大路径, 需要 左/右子树的最大路径(max), 左/右子树的以left/right开始的最大路径(size)。

var max int

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max = math.MinInt32
	process(root)
	return max - 1
}

func process(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := node.Left
	right := node.Right
	leftSize := process(left)
	rightSize := process(right)
	size := 1

	if left != nil && left.Val == node.Val {
		size = leftSize + 1
	}
	if right != nil && right.Val == node.Val {
		size = Max(size, rightSize+1)
	}
	max = Max(max, size)
	if left != nil && right != nil && left.Val == node.Val && right.Val == node.Val {
		max = Max(max, leftSize+rightSize+1)
	}
	return size
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
