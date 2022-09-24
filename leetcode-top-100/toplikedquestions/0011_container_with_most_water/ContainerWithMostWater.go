package _011_container_with_most_water

// 盛水最多的容器
// https://leetcode.cn/problems/container-with-most-water/
func maxArea(height []int) int {
	max := 0
	left := 0
	right := len(height) - 1
	for left < right {
		max = Max(max, Min(height[left], height[right])*(right-left))
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
