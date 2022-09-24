package _015_three_sum

import (
	"sort"
)

// 三数之和
// https://leetcode.cn/problems/3sum/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i-1] != nums[i] {
			nexts := twoSum(nums, i+1, -nums[i])
			for _, cur := range nexts {
				cur = append(cur, nums[i])
				result = append(result, cur)
			}
		}
	}
	return result
}

func twoSum(nums []int, begin, target int) [][]int {
	left := begin
	right := len(nums) - 1
	result := make([][]int, 0)
	for left < right {
		if nums[left]+nums[right] > target {
			right--
		} else if nums[left]+nums[right] < target {
			left++
		} else { // 需要判断 nums[left] != nums[left-1]
			if left == begin || nums[left-1] != nums[left] {
				cur := make([]int, 2)
				cur[0] = nums[left]
				cur[1] = nums[right]
				result = append(result, cur)
			}
			left++
		}
	}
	return result
}
