package _035_search_insert_position

// 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = left + ((right - left) >> 1)
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			break
		}
	}
	return TernaryInt(target <= nums[mid], mid, mid+1)
}

func TernaryInt(flag bool, a, b int) int {
	if flag {
		return a
	}
	return b
}
