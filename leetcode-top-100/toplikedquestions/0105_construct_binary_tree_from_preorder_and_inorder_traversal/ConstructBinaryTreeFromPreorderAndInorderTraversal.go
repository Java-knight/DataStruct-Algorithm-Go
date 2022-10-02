package _105_construct_binary_tree_from_preorder_and_inorder_traversal

// 从前序与中序遍历序列构造二叉树
// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

// 使用map表做一点加速
func buildTree(preorder []int, inorder []int) *TreeNode {
	cache := make(map[int]int, 0)
	for i, val := range inorder {
		cache[val] = i
	}
	return process(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, cache)
}

// 在 preorder[l1, r1]和inorder[l2, r2]上构造树
func process(preorder []int, l1, r1 int, inorder []int, l2, r2 int, cache map[int]int) *TreeNode {
	if l1 > r1 { // base case1 越界
		return nil
	}
	node := &TreeNode{preorder[l1], nil, nil}
	if l1 == r1 { // base case2 已经找完了
		return node
	}

	// base case3 inorder上还有值, 找下preorder[l1] = inorder[findIndex]
	// 在 inorder 上找到 preorder[l1] = inorder[findIndex]
	findIndex := l2
	for ; findIndex <= r2; findIndex++ {
		if preorder[l1] == inorder[findIndex] {
			break
		}
	}
	node.Left = process(preorder, l1+1, r1-r2+findIndex, inorder, l2, findIndex-1, cache)
	node.Right = process(preorder, r1-r2+findIndex+1, r1, inorder, findIndex+1, r2, cache)
	return node
}

//// 暴力递归
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	return process(preorder, 0, len(preorder) - 1, inorder, 0, len(inorder) - 1)
//}
//
//// 在 preorder[l1, r1]和inorder[l2, r2]上构造树
//func process(preorder []int, l1, r1 int, inorder []int, l2, r2 int) *TreeNode {
//	if l1 > r1 {   // base case1 越界
//		return nil
//	}
//	node := &TreeNode{preorder[l1], nil, nil}
//	if l1 == r1 {  // base case2 已经找完了
//		return node
//	}
//
//	// base case3 inorder上还有值, 找下preorder[l1] = inorder[findIndex]
//	// 在 inorder 上找到 preorder[l1] = inorder[findIndex]
//	findIndex := l2
//	for ; findIndex <= r2; findIndex++ {
//		if preorder[l1] == inorder[findIndex] {
//			break
//		}
//	}
//	node.Left = process(preorder, l1 + 1, r1 - r2 + findIndex, inorder, l2, findIndex - 1)
//	node.Right = process(preorder, r1 - r2 + findIndex + 1, r1, inorder, findIndex + 1, r2)
//	return node
//}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
