package _053_maximum_subarray

import "math"

// 最大子数组和
// https://leetcode.cn/problems/maximum-subarray/

// 动态规划(空间优化)
func maxSubArray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	max := math.MinInt32
	cur := 0
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		max = Max(max, cur)               // 收集结果
		cur = TernaryInt(cur < 0, 0, cur) // 我加了nums[i]位置的数还导致我变成了负数, 那我就不能加, 从0开始
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TernaryInt(flag bool, a, b int) int {
	if flag {
		return a
	} else {
		return b
	}
}

//// 方法1: 动态规划
//func maxSubArray(nums []int) int {
//	if nums == nil || len(nums) == 0 {
//		return -1
//	}
//	dp := make([]int, len(nums))
//	dp[0] = nums[0]
//	max := dp[0]
//	for i := 1; i < len(nums); i++ {
//		case1 := nums[i]  // base case1 只有当前i位置的数
//		case2 := nums[i] + dp[i - 1]  // base case2 当前i位置的数+i-1位置上的最大值
//		dp[i] = Max(case1, case2)
//		max = Max(max, dp[i])  // 收集最优解的
//	}
//	return max
//}
