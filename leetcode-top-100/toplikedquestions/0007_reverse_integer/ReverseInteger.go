package main

// 整数反转
// https://leetcode.cn/problems/reverse-integer/

// (1) 基本的操作 res = res * 10 + x % 10; x /= 10
// (2) 将x转为负数操作, 因为负数的域大于正数
// (3) 防止出现越界
func reverse(x int) int {
	result := 0
	for x != 0 {
		// 判断是否越界
		if tmp := int32(result); (tmp*10)/10 != tmp {
			return 0
		}
		result = result*10 + x%10
		x /= 10
	}
	return result
}

func main() {

}
