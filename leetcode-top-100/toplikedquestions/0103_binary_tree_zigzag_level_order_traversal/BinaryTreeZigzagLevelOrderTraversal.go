package _103_binary_tree_zigzag_level_order_traversal

// 二叉树的锯齿层序遍历
// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/

// 思路(画图): level从第0层开始, root直接加入;
//      level到第1层, 先加左孩子, 再加右孩子, 头进尾出;
//      level到第2层, 先加右孩子, 再加左孩子, 尾进头出; 接下来同理
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := make([]*TreeNode, 0)
	flag := true
	queue = append(queue, root)
	size := 1
	for len(queue) != 0 {
		size = len(queue)
		cur := make([]int, size)
		if flag { // 先左后右, 队尾进队头出
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
		} else { // 先右后左, 队头进队尾出
			for i := 0; i < size; i++ {
				node := queue[len(queue)-1]
				queue = queue[:len(queue)-1]
				cur[i] = node.Val
				if node.Right != nil {
					queue = append([]*TreeNode{node.Right}, queue...)
				}
				if node.Left != nil {
					queue = append([]*TreeNode{node.Left}, queue...)
				}
			}
		}
		flag = !flag
		result = append(result, cur)
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
