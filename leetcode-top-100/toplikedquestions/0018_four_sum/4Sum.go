package _018_four_sum

import (
	"sort"
)

// 四数之和
// https://leetcode.cn/problems/4sum/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(nums, 0, target, 4)
}

// 求nSum之和, nums必须是有序的, begin 是开始位置, n是当前求的n个数之和为target
// 递归调用, 只有在twoSum的时候需要使用双指针
func nSum(nums []int, begin, target, n int) [][]int {
	if n == 2 {
		return twoSum(nums, begin, target)
	}
	result := make([][]int, 0)
	for i := begin; i < len(nums)-n+1; i++ {
		if i == begin || nums[i-1] != nums[i] {
			if isSubtractOverflow(target, nums[i]) {
				return result
			}
			nexts := nSum(nums, i+1, target-nums[i], n-1)
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

// 判断是否发生溢出, 仿照Math.subtractExact()改造的
func isSubtractOverflow(x, y int) bool {
	r := x - y
	if ((x ^ y) & (x ^ r)) < 0 {
		return true
	}
	return false
}
