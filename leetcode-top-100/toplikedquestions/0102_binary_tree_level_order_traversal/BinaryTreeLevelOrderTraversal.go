package _102_binary_tree_level_order_traversal

// 二叉树的层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	var size int
	queue = append(queue, root)

	for len(queue) != 0 {
		size = len(queue)
		cur := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			cur[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, cur)
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
