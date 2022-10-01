package _098_validate_binary_search_tree

import "math"

// 验证二叉搜索树
// https://leetcode.cn/problems/validate-binary-search-tree/

// 方法3: Morris遍历(时间复杂度O(N), 空间复杂度O(1))—>空间准确来说应该是O(2N)
func isValidBST(root *TreeNode) bool {
	var mostRight *TreeNode
	cur := root
	pre := math.MinInt64
	flag := true
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil { // 第一次访问(连线)
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else { // 第二次访问(断线)
				mostRight.Right = nil
			}
		}
		if pre >= cur.Val {
			flag = false
		}
		pre = cur.Val
		cur = cur.Right
	}
	return flag
}

//// 方法1: 递归版本(时间复杂度O(N), 空间复杂度O(H))
//func isValidBST(root *TreeNode) bool {
//	return process(root, math.MinInt64, math.MaxInt64)
//}
//
//func process(node *TreeNode, min, max int64) bool {
//	if node == nil {
//		return true
//	}
//	if int64(node.Val) >= max || int64(node.Val) <= min {
//		return false
//	}
//	return process(node.Left, min, int64(node.Val)) && process(node.Right, int64(node.Val), max)
//}

//// 方法2: 迭代版本(时间复杂度O(N), 空间复杂度O(H))
//func isValidBST(root *TreeNode) bool {
//	pre := math.MinInt64
//	var stack []interface{}
//	cur := root
//	for len(stack) != 0 || cur != nil {
//		if cur != nil {
//			stack = append(stack, cur)
//			cur = cur.Left
//		} else {
//			cur = stack[len(stack)-1].(*TreeNode)
//			stack = stack[:len(stack) - 1]
//			if pre >= cur.Val {
//				return false
//			}
//			pre = cur.Val
//			cur = cur.Right
//		}
//	}
//	return true
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
