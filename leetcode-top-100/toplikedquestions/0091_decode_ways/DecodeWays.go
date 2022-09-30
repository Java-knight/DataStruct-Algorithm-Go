package _091_decode_ways

// 解码方法
// https://leetcode.cn/problems/decode-ways/

// 动态规划(从左往右的模型)
func numDecodings(s string) int {
	if s == "" || len(s) == 0 {
		return 0
	}
	size := len(s)
	// dp[index]: 就str[index....] 能转换多少有效, 返回方法数
	dp := make([]int, size+1)
	dp[size] = 1
	// dp[i] 依赖dp[i + 1]和dp[i + 2]位置(从右向左的尝试模型)
	for i := size - 1; i >= 0; i-- {
		if s[i] != '0' {
			// base case1 i位置是1~9
			dp[i] = dp[i+1]
			if i+1 == size {
				continue
			}
			// base case2 (i,i+1) 位置(小于26) "17"
			num := (s[i]-'0')*10 + (s[i+1] - '0')
			if num <= 26 {
				dp[i] += dp[i+2]
			}
		}
	}
	return dp[0]
}

//// 暴力递归(从右往左的尝试模型)
//func numDecodings(s string) int {
//	if s == "" || len(s) == 0 {
//		return 0
//	}
//
//	return process(s, 0)
//}
//
//// str[index....] 能转换多少有效, 返回方法数
//func process(str string, index int) int {
//	if index == len(str) {
//		return 1
//	}
//	// index == 0, 是不能转的; index(1~9)
//	if str[index] == '0' {
//		return 0
//	}
//	// base case1 str[index]属于1~9
//	ways := process(str, index+1)
//	if index + 1 == len(str) {
//		return ways
//	}
//
//	// base case2 (index index+1) 一起去看(必须和起来<=26) "16"
//	num := (str[index]-'0')*10 + (str[index+1] - '0')
//	if num <= 26 {
//		ways += process(str, index + 2)
//	}
//	return ways
//}
