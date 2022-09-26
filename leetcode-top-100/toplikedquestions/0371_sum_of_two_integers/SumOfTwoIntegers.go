package _371_sum_of_two_integers

// 29是除法
// 两整数之和
// https://leetcode.cn/problems/sum-of-two-integers/

// a ^ b ——> a+b的无进位信息; (a & b) << 1 ——> a+b的进位信息
// a+b只要 将无进位信息赋值给a, 进位信息赋值给b, 再次相加。直到将进位信息消耗完
// 比如：a=10   01010   00101    10001    11001
//      b=15   01111   10100    01000    00000
// a+b = 25  ——> 11001
func getSum(a int, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b      // 无进位信息
		b = (a & b) << 1 // 进位信息
		a = sum
	}
	return sum
}

// 获取一个数的相反数, 不能用负号(go语言的取反符号是^, 而Java是~)
// 10    1010    1111...0101    1111...0110
// -10: 二进制是10的二进制取反+1    1111...0110
func negNum(n int) int {
	return getSum(^n, 1)
}

// 减法
// a - b == a + (-b)
func minus(a, b int) int {
	return getSum(a, negNum(b))
}
