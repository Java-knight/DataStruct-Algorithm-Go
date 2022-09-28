package _039_combination_sum

import (
	"sort"
)

// 组合总和
// https://leetcode.cn/problems/combination-sum/

var result [][]int // 结果
var path []int     // 每一条路径
func combinationSum(candidates []int, target int) [][]int {
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
			if target >= candidates[i] {
				// 标记现场
				path = append(path, candidates[i])
				// 进入下层决策
				process(candidates, target-candidates[i], i)
				// 恢复现场
				path = path[:len(path)-1]
			}
		}
	}
}
