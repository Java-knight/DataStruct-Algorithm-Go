package _075_sort_colors

// 颜色分类
// https://leetcode.cn/problems/sort-colors/

// 荷兰国旗问题(让arr按num划分成小于区、等于区、大于区)
// 思路: 搞左右两个指针(都是边界指针), 比如left=-1, right=arr.length
// left推着index向右跑, right向左跑来
func sortColors(nums []int) {
	left := -1
	right := len(nums)
	index := 0
	for index < right {
		if nums[index] == 1 { // 等于区
			index++
		} else if nums[index] == 0 { // 小于区
			left++
			nums[index], nums[left] = nums[left], nums[index]
			index++
		} else { // 大于区
			right--
			nums[index], nums[right] = nums[right], nums[index]
		}
	}
}
