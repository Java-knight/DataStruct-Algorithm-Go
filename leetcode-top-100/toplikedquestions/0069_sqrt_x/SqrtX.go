package _069_sqrt_x

// x的平方根
// https://leetcode.cn/problems/sqrtx/

// x不可能输入负数
// 二分法, 每次都看mid*mid和x的关系, mid*mid > x, right向左移动; mid*mid <= x, left向右移动
// 使用int64去接变量, 防止溢出
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 4 {
		return 1
	}
	result := 1
	var left, right, mid int64 = 0, int64(x), 0
	for left <= right {
		mid = left + ((right - left) >> 1)
		if mid*mid <= int64(x) {
			result = int(mid)
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
