package _094_binary_tree_inorder_traversal

// 二叉树的中序遍历
// https://leetcode.cn/problems/binary-tree-inorder-traversal/

// 方法3: Morris遍历(时间复杂度O(N), 空间复杂度O(1))—>空间准确来说应该是O(2N)
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	cur := root
	var mostRight *TreeNode
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur // 第一次访问(连线)
				cur = cur.Left
				continue
			} else { // mostRight.Right == cur
				mostRight.Right = nil // 第二次访问(断线)
			}
		}
		result = append(result, cur.Val)
		cur = cur.Right
	}
	return result
}

//// 方法1: 递归版本(时间复杂度O(N), 空间复杂度O(H))
//var result []int
//func inorderTraversal(root *TreeNode) []int {
//	result = make([]int, 0)
//
//	process(root)
//	return result
//}
//
//func process(node *TreeNode) {
//	if node == nil {
//		return
//	}
//	process(node.Left)
//	result = append(result, node.Val)
//	process(node.Right)
//}

//// 方法2: 迭代版本(时间复杂度O(N), 空间复杂度O(H))
//func inorderTraversal(root *TreeNode) []int {
//	result := make([]int, 0)
//	var stack []interface{}
//	cur := root
//	for len(stack) != 0 || cur != nil {
//		if cur != nil {
//			stack = append(stack, cur)
//			cur = cur.Left
//		} else {
//			cur = stack[len(stack) - 1].(*TreeNode)
//			stack = stack[:len(stack) - 1]
//			result = append(result, cur.Val)
//			cur = cur.Right
//		}
//	}
//	return result
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
