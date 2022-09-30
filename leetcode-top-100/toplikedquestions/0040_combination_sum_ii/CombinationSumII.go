package _040_combination_sum_ii

import (
	"sort"
)

// 组合总和II
// https://leetcode.cn/problems/combination-sum-ii/

var result [][]int // 结果
var path []int     // 每一条路径
func combinationSum2(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	path = make([]int, 0)
	sort.Ints(candidates)
	process(candidates, target, 0)
	return result
}

func process(candidates []int, target int, index int) {
	if target == 0 { // 收集答案
		cur := make([]int, len(path))
		copy(cur, path)
		result = append(result, cur)
	} else {
		for i := index; i < len(candidates); i++ {
			if target < candidates[i] { // 剪支优化
				break
			}
			if i > index && candidates[i] == candidates[i-1] { // 去重
				continue
			}
			// 标记现场
			path = append(path, candidates[i])
			// 进入下层决策
			process(candidates, target-candidates[i], i+1)
			// 恢复现场
			path = path[:len(path)-1]
		}
	}
}
