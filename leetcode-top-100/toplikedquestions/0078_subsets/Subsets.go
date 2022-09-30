package _078_subsets

// 子集
// https://leetcode.cn/problems/subsets/

// 思路: 一个简单的暴力递归, 每个数分为要 和 不要, 叶子节点收集答案
var result [][]int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	path := make([]int, 0)
	process(path, 0, nums)
	return result
}

func process(path []int, index int, nums []int) {
	if index == len(nums) { // 收集答案
		ans := make([]int, len(path))
		copy(ans, path)
		result = append(result, ans)
	} else {
		// 【不要index位置的数】
		// 标记现场(不要当前这个数, 所以不需要标记现场)
		// 进入下层决策
		process(path, index+1, nums)
		// 恢复现场(因为没有标记现场, 也不需要恢复)

		// 【要index位置的数】
		// 标记现场(要当前这个数, 所以需要标记现场)
		path = append(path, nums[index])
		// 进入下层决策
		process(path, index+1, nums)
		// 恢复现场
		path = path[:len(path)-1]
	}
}
