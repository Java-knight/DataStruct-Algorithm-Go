package _034_find_first_and_last_position_of_element_in_sorted_array

// 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/

// 二分找最右边和最左边, 分两个函数调
func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if nums == nil || len(nums) == 0 {
		return result
	}
	result[0] = findLeft(nums, target)
	result[1] = findRight(nums, target)
	return result
}

func findLeft(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	leftIndex := -1
	for left <= right {
		mid = left + ((right - left) >> 1)
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			leftIndex = mid
			right = mid - 1
		}
	}
	return leftIndex
}

func findRight(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	rightIndex := -1
	for left <= right {
		mid = left + ((right - left) >> 1)
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			rightIndex = mid
			left = mid + 1
		}
	}
	return rightIndex
}
