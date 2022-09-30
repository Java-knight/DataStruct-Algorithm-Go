package _047_permutations_ii

// 全排列II
// https://leetcode.cn/problems/permutations-ii/

var result [][]int

func permuteUnique(nums []int) [][]int {
	result = make([][]int, 0)
	process(0, nums)
	return result
}

func process(index int, nums []int) {
	if index == len(nums) { // 收集答案
		cur := make([]int, len(nums))
		copy(cur, nums)
		result = append(result, cur)
	} else {
		// 本次path中重复的数字加入到set中
		set := make(map[int]int)
		for i := index; i < len(nums); i++ {
			if _, ok := set[nums[i]]; !ok {
				set[nums[i]] = -1
				// 标记现场
				nums[i], nums[index] = nums[index], nums[i]
				// 进入下层决策
				process(index+1, nums)
				// 恢复现场
				nums[i], nums[index] = nums[index], nums[i]
			}
		}
	}
}
