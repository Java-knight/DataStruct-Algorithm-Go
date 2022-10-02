package _106_construct_binary_tree_from_inorder_and_postorder_traversal

// 从中序和后续遍历构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

// 使用map表做一点加速
func buildTree(inorder []int, postorder []int) *TreeNode {
	cache := make(map[int]int, 0)
	for i := range inorder {
		cache[inorder[i]] = i
	}
	return process(postorder, 0, len(postorder)-1, inorder, 0, len(inorder)-1, cache)
}

func process(postorder []int, l1, r1 int, inorder []int, l2, r2 int, cache map[int]int) *TreeNode {
	if l1 > r1 { // base case1 叶子节点的左右孩子
		return nil
	}
	node := &TreeNode{postorder[r1], nil, nil}
	if l1 == r1 { // base case2 叶子节点
		return node
	}

	// base case3 非叶子节点, 在inorder[findIndex] == postorder[r1]
	findIndex := cache[postorder[r1]]
	node.Left = process(postorder, l1, l1+findIndex-l2-1, inorder, l2, findIndex-1, cache)
	node.Right = process(postorder, l1+findIndex-l2, r1-1, inorder, findIndex+1, r2, cache)
	return node
}

//// 暴力递归
//func buildTree(inorder []int, postorder []int) *TreeNode {
//	return process(postorder, 0, len(postorder) - 1, inorder, 0, len(inorder) - 1)
//}
//
//func process(postorder []int, l1, r1 int, inorder []int, l2, r2 int) *TreeNode {
//	if l1 > r1 {  // base case1 叶子节点的左右孩子
//		return nil
//	}
//	node := &TreeNode{postorder[r1], nil, nil}
//	if l1 == r1 {  // base case2 叶子节点
//		return node
//	}
//
//	// base case3 非叶子节点, 在inorder[findIndex] == postorder[r1]
//	findIndex := l2
//	for findIndex < r2 {
//		if postorder[r1] == inorder[findIndex] {
//			break
//		}
//		findIndex++
//	}
//	node.Left = process(postorder,l1, l1 + findIndex - l2 - 1, inorder, l2, findIndex - 1);
//	node.Right = process(postorder, l1 + findIndex - l2, r1 - 1, inorder, findIndex + 1, r2);
//	return node
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
