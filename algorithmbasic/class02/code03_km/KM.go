package main

// 题目：一个数组中有一种数出现K次，其它数都出现了M次，M>1，K<M。找到出现了K次的数。
// 要求，额外空间复杂度O(1)，时间复杂度O(N)

// 解法1: 最简洁的方法
func km(arr []int, k, m int) int {
	help := make([]int, 32)
	for _, val := range arr {
		for i := 0; i < 32; i++ {
			help[i] += (val >> i) & 1
		}
	}

	result := 0
	for i := 0; i < 32; i++ {
		help[i] %= m
		if help[i] != 0 {
			result |= 1 << i
		}
	}
	return result
}

func main() {

}
