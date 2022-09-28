//package _046_permutations
package main

func main() {
	arr := []int{1, 2, 3}
	permute(arr)
}

// 全排列
// https://leetcode.cn/problems/permutations/

var result [][]int

func permute(nums []int) [][]int {
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
		for i := index; i < len(nums); i++ {
			// 保存现场
			nums[i], nums[index] = nums[index], nums[i]
			// 进入下层决策
			process(index+1, nums)
			// 恢复现场
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
}
