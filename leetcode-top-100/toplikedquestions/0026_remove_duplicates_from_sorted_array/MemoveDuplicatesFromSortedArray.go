package _026_remove_duplicates_from_sorted_array

// 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	if len(nums) < 2 {
		return len(nums)
	}
	result := 0
	for index := range nums {
		if index == 0 || nums[index-1] != nums[index] {
			nums[result] = nums[index]
			result++
		}
	}
	return result
}
