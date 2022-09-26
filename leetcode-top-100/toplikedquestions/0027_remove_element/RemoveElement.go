package _027_remove_element

// 移除元素
// https://leetcode.cn/problems/remove-element/

// 双指针(左右指针, 重复元素避免的大量赋值操作)
func removeElement(nums []int, val int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}

//// 快慢指针(遇到重复元素需要不断的赋值)
//func removeElement(nums []int, val int) int {
//	if nums == nil || len(nums) == 0 {
//		return 0
//	}
//
//	slow := 0
//	for fast := 0; fast < len(nums); fast++ {
//		if nums[fast] != val {
//			nums[slow] = nums[fast]
//			slow++
//		}
//	}
//	return slow
//}
