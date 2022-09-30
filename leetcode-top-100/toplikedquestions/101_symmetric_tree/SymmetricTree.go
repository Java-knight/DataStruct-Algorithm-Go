package _01_symmetric_tree

// 对称二叉树
// https://leetcode.cn/problems/symmetric-tree/

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

// 一棵树是原始树   node1
// 另一棵树是翻面树 node2
func isMirror(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 != nil && node2 != nil {
		return node1.Val == node2.Val &&
			isMirror(node1.Left, node2.Right) &&
			isMirror(node2.Right, node1.Left)
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
