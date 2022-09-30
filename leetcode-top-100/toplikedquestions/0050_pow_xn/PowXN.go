package _050_pow_xn

import "math"

// Pow(x, n)
// https://leetcode.cn/problems/powx-n/

// n是正数: x^n ——> n = 二进制, 比如x^75 = x^64 *x^8 *x^2 *x^1
// n是负数: x^n —> -n = 二进制, 结果 1/result
// 注意: Integer.MIN_VALUE系统最小是无法转为正数的(可以通过+1 就能转为正数)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == math.MinInt32 { // num^min ——> 1/num^max = 0
		return TernaryDouble(x == 1 || x == -1, 1, 0)
	}
	// 当n = Integer.MIN_VALUE, 让它+1 就可以转为正数了, 后面结果*x就ok了
	pow := int32(math.Abs(float64(TernaryInt(n == math.MinInt32, n+1, n))))
	var result float64 = 1
	cur := x // 2^1、2^2、2^3、2^4...

	for pow != 0 {
		if pow&1 != 0 { // pow该位置的二进制是1 收集结果
			result *= cur
		}
		pow >>= 1
		cur *= cur
	}
	if n == math.MinInt32 {
		result *= x
	}
	return TernaryDouble(n < 0, 1/result, result)
}

func TernaryDouble(flag bool, a, b float64) float64 {
	if flag {
		return a
	} else {
		return b
	}
}

func TernaryInt(flag bool, a, b int) int {
	if flag {
		return a
	} else {
		return b
	}
}

//func myPow(x float64, n int) float64 {
//	if n == 0 {
//		return 1
//	}
//	if n == math.MinInt32 { // num^min ——> 1/num^max = 0
//		return TernaryDouble(x == 1 || x == -1, 1, 0)
//	}
//	var result float64 = 1
//	cur := x // 2^1、2^2、2^3、2^4...
//	pow := int32(math.Abs(float64(n)))
//	for pow != 0 {
//		if pow&1 != 0 { // pow该位置的二进制是1 收集结果
//			result *= cur
//		}
//		pow >>= 1
//		cur *= cur
//	}
//	return TernaryDouble(n < 0, 1/result, result)
//}
