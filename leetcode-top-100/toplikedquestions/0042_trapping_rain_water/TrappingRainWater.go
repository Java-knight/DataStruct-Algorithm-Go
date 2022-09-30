package _042_trapping_rain_water

// 接雨水
// https://leetcode.cn/problems/trapping-rain-water/

// 思路: i位置的水 = min(max[0, i-1], max[i+1, size-1]) - i
// 使用两个指针就可以, left和right指针, 记录left指针左边的最大值leftMax, 记录right指针右边的做大值rightMax
// 每次leftMax和rightMax进行比较, 谁小计算谁。
// 注意: 0和size-1位置必不可能储水
func trap(height []int) int {
	if height == nil || len(height) < 3 {
		return 0
	}
	result := 0
	size := len(height)
	left := 1
	right := size - 1
	leftMax := height[0]
	rightMax := height[size-1]
	for left <= right {
		if leftMax <= rightMax {
			result += max(0, leftMax-height[left]) // 因为height[left]可能比leftMax大(出现负数), 所以需要取最大值
			leftMax = max(leftMax, height[left])
			left++
		} else {
			result += max(0, rightMax-height[right]) // 因为height[right]可能比rightMax大(出现负数), 所以需要取最大值
			rightMax = max(rightMax, height[right])
			right--
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//// 方法1: 两个辅助数组, 有点浪费空间
//// 从左到右搞一个数组, 数组的元素对应从左到右当前index位置的最大值
//// 从右到左搞一个数组, 数组的元素对应从右到左当前index位置的最大值
//// arr   [3, 2, 5, 2, 6, 2, 2, 9, 2]
//// left  [3, 3, 5, 5, 6, 6, 6, 9, 9]
//// right [9, 9, 9, 9, 9, 9, 9, 9, 2]
//func trap(height []int) int {
//	if height == nil || len(height) < 3 {
//		return 0
//	}
//	result := 0
//	size := len(height)
//	leftMax := make([]int, size)
//	rightMax := make([]int, size)
//	leftMax[0] = height[0]
//	rightMax[size - 1] = height[size - 1]
//	// leftMax数组
//	for i := 1; i < size; i++ {
//		if leftMax[i - 1] < height[i] {
//			leftMax[i] = height[i]
//		} else {
//			leftMax[i] = leftMax[i - 1]
//		}
//	}
//	// rightMax数组
//	for i := size - 2; i >= 0; i-- {
//		if rightMax[i + 1] < height[i] {
//			rightMax[i] = height[i]
//		} else {
//			rightMax[i] = rightMax[i + 1]
//		}
//	}
//
//	// 收集结果
//	for i := 1; i < size - 1; i++ {
//		result += min(leftMax[i], rightMax[i]) - height[i]
//	}
//	return result
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}
