package _067_add_binary

import "strconv"

// 二进制求和
// https://leetcode.cn/problems/add-binary/

func addBinary(a string, b string) string {
	result := ""
	carry := 0 // 充当sum和carry的角色
	size := Max(len(a), len(b))
	for i := 0; i < size; i++ {
		if i < len(a) { // 倒叙进行
			carry += int(a[len(a)-1-i] - '0')
		}
		if i < len(b) {
			carry += int(b[len(b)-1-i] - '0')
		}
		result = strconv.Itoa(carry%2) + result
		carry >>= 1
	}
	if carry == 1 {
		result = strconv.Itoa(carry) + result
	}
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
