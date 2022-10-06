//package _129_sum_root_to_leaf_numbers

package main

import (
	"log"
	"math"
)

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	log.Println(sumNumbers(root))
}

// 求根节点到叶节点数字之和
// https://leetcode.cn/problems/sum-root-to-leaf-numbers/

// 思路: DFS
var result int

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result = 0
	path := make([]int, 0)
	path = append(path, root.Val)
	process(path, root)
	return result
}

func process(path []int, node *TreeNode) {
	if node.Left == nil && node.Right == nil { // 收集答案
		result += conversion(path)
	} else {
		if node.Left != nil {
			// 标记现场(左)
			path = append(path, node.Left.Val)
			// 进入下层决策(左)
			process(path, node.Left)
			// 恢复现场(左)
			path = path[:len(path)-1]
		}
		if node.Right != nil {
			// 标记现场(右)
			path = append(path, node.Right.Val)
			// 进入下层决策(右)
			process(path, node.Right)
			// 恢复现场(右)
			path = path[:len(path)-1]
		}
	}
}

func conversion(path []int) int {
	cur := 0
	i := len(path) - 1
	j := 0
	for i >= 0 {
		cur += path[i] * int(math.Pow(10, float64(j)))
		i--
		j++
	}
	return cur
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
