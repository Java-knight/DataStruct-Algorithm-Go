package main

import "log"

// 题目: 一个数组中有一种数出现了奇数次，其它数都出现了偶数次，怎么找到并打印这种数
// 题目: 怎么把一个int类型的数，提取出最右侧的1来
// 题目: 一个数组种有两种数出现了奇数次，其它数都出现了偶数次，怎么找到并打印这两种数

// arr 中, 只有一种数, 出现奇数次
func printOddTimesNum1(arr []int) {
	eor := 0
	for _, val := range arr {
		eor ^= val
	}

	log.Println(eor)
}

// arr 中, 有两种数, 出现奇数次
// (1) 先找出两个奇数 a ^ b = eor
// (2) 找出 eor 左边第一个 1（其实哪个1都可以，因为这就是a和b的不同之处）  rightOne = eor ^ (-eor)
// (3) 去找 arr 中这个位置 != 0的数 if ((rightOne & arr[i]) != 0) { onlyOne ^= arr[i]} 表示用其中一个奇数a/b ^ arr[i](出现偶数次)
// (4) onlyOne 就是一个奇数, 另一个就是 onlyOne ^ eor ——> a/b ^ a ^ b = a/b
func printOddTimesNum2(arr []int) {
	eor := 0
	for _, val := range arr {
		eor ^= val
	}

	// a 和 b是两个数
	// eor != 0
	// eor 最右侧的1, 提取出来
	// eor:      0110 0101 1011 1000
	// rightOne: 0000 0000 0000 1000

	// -eor:     1001 1010 0100 1000
	// rightOne = eor & (-eor)
	// 结论: 取反+1，刚好是把最后一位变成了和原码相同的，其它位数都相反
	rightOne := eor & (-eor)
	onlyOne := 0
	for _, val := range arr {
		if (rightOne & val) != 0 {
			onlyOne ^= val
		}
	}
	log.Println(onlyOne, ", ", (onlyOne ^ eor))
}

func main() {
	arr := []int{6, 10, 6, 6, 4, 4, 12, 12, 12, 12, 3, 3}
	printOddTimesNum2(arr)
}
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		rightOne := num & (-num)
		num ^= rightOne
		count++
	}
	return count
}
