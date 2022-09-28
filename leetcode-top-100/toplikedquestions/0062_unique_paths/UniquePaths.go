package _062_unique_paths

// 不同路径
// https://leetcode.cn/problems/unique-paths/

// question: 从左上角走到右下角, 有多少种不同的路径?
// 思路: 每次只能向下或者向右, 那么不就是排列组合嘛, 直接C(l+r, l/r)
//      C(N, M) = N!/(M!*(N-M)!) = C(N, (N-M))
// 注意: 怎么计算阶乘尽量不越界, 相乘前进行约分(需要求最大公约数, 用long放阶乘的结果)
func uniquePaths(m int, n int) int {
	// C(all, part)
	part := m - 1
	all := n + m - 2
	var o1 int64 = 1
	var o2 int64 = 1

	i := part + 1
	j := 1
	// C(all, part) = all!/(part!*(all-part)!) = ((part+1)*(part+2)*...all)/(1*2*...(all-part))
	for i <= all || j <= all-part {
		o1 *= int64(i)
		o2 *= int64(j)
		gcd := Gcd(o1, o2)
		o1 /= gcd
		o2 /= gcd

		i++
		j++
	}
	return int(o1) // 最终o2肯定是被约分完了
}

// Gcd greatest common divisor 最大公约数
func Gcd(m, n int64) int64 {
	if n == 0 {
		return m
	} else {
		return Gcd(n, m%n)
	}
}
