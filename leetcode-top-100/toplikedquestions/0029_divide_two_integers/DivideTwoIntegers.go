package _029_divide_two_integers

import (
	"math"
)

// 371 是加法
// 两数相除
// https://leetcode.cn/problems/divide-two-integers/

// 除法
// 边界条件: 系统最小表示 math.MinInt32
// (1) dividend 和 divisor 都是系统最小, 结果 1
// (2) dividend不是系统最小, divisor 是系统最小,结果0. 因为结果范围一定是(-1, 1)
// (3) dividend是系统最小, divisor 是-1, 结果系统最大, 因为系统最小的绝对值比系统最大放不下;
//     divisor 不是系统最小也不是-1, 需要使用补偿法, 假设系统最大是-20, divisor是4,
//     我就先给dividend+1 = -19, res = -19/4 = -4, (dividend-res*divisor)/divisor
//     需要补偿的值 = (-20-(-4)*4)/4 = -1, 结果 res+补偿值 = -4+(-1) = -5
// (4) dividend 和 divisor 都不是系统最小, 直接调用div进行计算(div会将dividend和divisor使用它们的绝对值进行计算)
// 问题: 为啥(3) 不能直接使用(4)呢? 要用什么补偿法, 因为(4) 需要转换为绝对值, (3)中的系统最小值直接就溢出了
func divide(dividend int, divisor int) int {
	// base case1/2
	if divisor == math.MinInt32 {
		return TernaryInt(dividend == math.MinInt32, 1, 0)
	}
	// base case3
	if dividend == math.MinInt32 {
		if divisor == negNum(1) { // divisor == -1
			return math.MaxInt32
		}
		result := div(add(dividend, 1), divisor)
		// res + (dividend-res*divisor)/divisor
		return add(result, div(minus(dividend, multi(result, uint(divisor))), divisor))
	}
	// base case4
	return div(dividend, divisor)
}

// 除法: 将a和b转化为绝对值, 使用a除b
// 思路: 16*12 = 16*2 + 16*10 = 16*2^2 + 16*2^3 = 192
// 将乘法反者来, 192/12 = 192-2^3*12-2^2*12-2^1*12-2^0*12 = 32
// 192    11000000    1100000    110000    11000    1100=32
// 12      1100000     110000     11000     1100    1100  (res就是12每次左移的位数)
// res  =    2^3   +    2^2    +    2^1   +   2^0  (代码中使用的是|符号作为+)
func div(a, b int) int {
	x := TernaryInt(isNeg(a), negNum(a), a)
	y := TernaryInt(isNeg(b), negNum(b), b)
	result := 0
	// i = 31; i > -1; i--
	for i := 31; i > negNum(1); i = minus(i, 1) {
		if (x >> i) >= y {
			result |= (1 << i)
			x = minus(x, y<<i)
		}
	}
	return TernaryInt(XORBool(isNeg(a), isNeg(b)), negNum(result), result)
}

// 乘法
// 15 * 10: 一般是用0*15+10*25
// 代码实现一样的, 只是把15 和 10转成二进制进行运算
// 15    1111    11110    111100    1111000
// 10    1010      101        10          1
// res =  0    +   30   +    0     +   80    = 150
func multi(a int, b uint) int {
	result := 0
	for b != 0 {
		if (b & 1) != 0 {
			result = add(result, a)
		}
		a <<= 1
		b >>= 1
	}
	return result
}

// a ^ b ——> a+b的无进位信息; (a & b) << 1 ——> a+b的进位信息
// a+b只要 将无进位信息赋值给a, 进位信息赋值给b, 再次相加。直到将进位信息消耗完
// 比如：a=10   01010   00101    10001    11001
//      b=15   01111   10100    01000    00000
// a+b = 25  ——> 11001
func add(a int, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b      // 无进位信息
		b = (a & b) << 1 // 进位信息
		a = sum
	}
	return sum
}

// 减法
// a - b == a + (-b)
func minus(a, b int) int {
	return add(a, negNum(b))
}

// 是否是负数
func isNeg(n int) bool {
	return n < 0
}

// 获取一个数的相反数, 不能用负号(go语言的取反符号是^, 而Java是~)
// 10    1010    1111...0101    1111...0110
// -10: 二进制是10的二进制取反+1    1111...0110
func negNum(n int) int {
	return add(^n, 1)
}

func TernaryInt(flag bool, a, b int) int {
	if flag {
		return a
	}
	return b
}

// 布尔类型的异或操作
func XORBool(a, b bool) bool {
	return (a || b) && !(a && b)
}
