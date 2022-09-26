package _033_search_in_rotataed_sorted_array

// 搜索旋转排序数组
// https://leetcode.cn/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = left + ((right - left) >> 1)
		// base case1 二分直接找到
		if nums[mid] == target {
			return mid
		}
		// base case2 nums[left] = nums[mid] = nums[right], 不能使用二分了, 需要left从左到右找
		if nums[left] == nums[mid] && nums[left] == nums[right] {
			// (1) left从左到右依次追mid
			for left != mid && nums[left] == nums[mid] {
				left++
			}
			// (2) 当left追上mid, 就可以跳出继续二分了
			if left == mid {
				left = mid + 1
				continue
			}
		}
		// base case3 nums[left]、nums[mid]、nums[right] 出现不相等的情况
		if nums[left] != nums[mid] {
			if nums[mid] > nums[left] { // (1) left... mid, 有序二分
				if target >= nums[left] && target < nums[mid] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			} else { // (2) nums[mid] < nums[left]: left...mid, 发生了反转, left~x-1有序, x~mid有序; mid...right一定有序使用二分
				if target > nums[mid] && target <= nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
		} else { // base case4 nums[left] == nums[mid] ——> nums[mid] != nums[right]
			if nums[mid] < nums[right] { // (1) mid...right 有序二分
				if target > nums[mid] && target <= nums[right] {
					left = mid + 1
				} else {
					right = mid - 1
				}
			} else { // (2) nums[mid] > nums[right]: mid...right 发生了反转, mid~x-1, x~right 有序, left...mid有序二分
				if target > nums[left] && target < nums[mid] {
					right = mid - 1
				} else {
					left = mid + 1
				}
			}
		}
	}
	return -1
}
