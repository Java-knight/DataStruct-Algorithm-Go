package _066_plus_one

// 加一
// https://leetcode.cn/problems/plus-one/

func plusOne(digits []int) []int {
	size := len(digits)
	for i := size - 1; i >= 0; i-- { // base case1 不是全9的
		if digits[i] < 9 { // 小于0, 直接+1返回结果
			digits[i]++
			return digits
		}
		digits[i] = 0 // 当前位置是9 +1后, 将设位置改为0
	}
	// base case2 如果没有再 for 循环中return, 说明digits全是99...9
	result := make([]int, size+1)
	result[0] = 1
	return result
}
