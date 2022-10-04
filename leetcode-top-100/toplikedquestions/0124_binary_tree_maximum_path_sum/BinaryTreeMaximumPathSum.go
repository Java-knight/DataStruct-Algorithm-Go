package _124_binary_tree_maximum_path_sum

import (
	"math"
)

// 二叉树中的最大路径和
// https://leetcode.cn/problems/binary-tree-maximum-path-sum/

// 思路: 按照x进行划分, x就是树上的某个node
// base case1 x无关: (2种情况求max)
//     a. leftTree(不含x, x的左子树的最大路径和)
//     b. rightTree(不含x, x的右子树的最大路径和)
// base case2 x有关: (4种情况求max)
//     a. selfX的值(只包含x一个)
//     b. leftPath, 以x为头向左子树连接
//     c. rightPath, 以x为头向右子树连接
//     d. leftRightPath, 以x为头向左子树和右子树都连接
// result: 求base case1 和 base case2的max
// 代码实现: 设计一个Info{result, maxPathHead}结构,
// result就是6种情况的max, maxPathHead 就是base case2中a,b,c的max
// 注意: 为啥maxPathHead 不是没有算base case2 的d, 因为两边都连接相当于head已经不是x节点了
func maxPathSum(root *TreeNode) int {
	result, _, _ := process(root)
	return result
}

// flag表示是否走到了nil节点
func process(x *TreeNode) (result, maxPathHead int, flag bool) {
	if x == nil { // 遍历到空了
		return 0, 0, false
	}
	leftResult, leftHead, leftFlag := process(x.Left)
	rightResult, rightHead, rightFlag := process(x.Right)

	// base case1
	p1 := Ternary(leftFlag, Max(math.MinInt32, leftResult), math.MinInt32)
	p2 := Ternary(rightFlag, Max(math.MinInt32, rightResult), math.MinInt32)
	// base case2
	p3 := x.Val
	p4 := Ternary(leftFlag, x.Val+leftHead, math.MinInt32)
	p5 := Ternary(rightFlag, x.Val+rightHead, math.MinInt32)
	p6 := Ternary(leftFlag && rightFlag, x.Val+leftHead+rightHead, math.MinInt32)

	// result
	maxPathHead = Max(Max(p4, p5), p3)
	result = Max(Max(Max(p1, p2), p6), maxPathHead)
	return result, maxPathHead, true
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Ternary(flag bool, a, b int) int {
	if flag {
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
