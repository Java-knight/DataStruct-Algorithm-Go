package _107_binary_tree_level_order_traversal_ii

// 二叉树的层序遍历II
// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/

// 问题: 从底向上遍历
func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	var queue []*TreeNode
	queue = append(queue, root)
	size := 0
	for len(queue) != 0 {
		size = len(queue)
		cur := make([]int, 0)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			cur = append(cur, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append([][]int{cur}, result...) // 每次都头插进入, 最后得到的list结果就是最底向上
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
