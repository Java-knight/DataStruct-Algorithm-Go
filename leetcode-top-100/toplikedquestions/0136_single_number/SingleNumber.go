package _136_single_number

// 只出现一次的数字
// https://leetcode.cn/problems/single-number

func singleNumber(nums []int) int {
	result := 0
	for _, val := range nums {
		result ^= val
	}
	return result
}
