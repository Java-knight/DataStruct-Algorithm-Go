//package _131_palindrome_partitioning
package main

import "fmt"

func main() {
	fmt.Println(partition("aab"))
}

// 分隔回文串
// https://leetcode.cn/problems/palindrome-partitioning/

// 思路: 递归写法, 加入s = "abcabc"是否是回文
// 比如: "a", "bcabc";   "ab", "cabc";   "abc", "abc";
//      "abca", "bc";   "abcab", "c";   "abcabc", "";
// str[L~R]是否是回文: (1) str[L] == str[R] (2) str[L+1, R-1]是否回文串
// dp表的设计: 直接使用一张bool二维dp表就可以完成, dp[i][j]就表示str[L, R]是否是回文,
//           dp表的左下部分不用(i>j -> L>R), dp表的初始化dp[i][i+1] = str[i] == str[i+1]
//           其它就是dp[i][j] = dp[i+1][j-1] && dp[i][j]
var result [][]string

func partition(s string) [][]string {
	dp := getDP([]byte(s))
	result = make([][]string, 0)
	path := make([]string, 0)
	process(path, 0, s, dp)
	return result
}

func getDP(str []byte) [][]bool {
	size := len(str)
	dp := make([][]bool, size)
	for i := range dp {
		dp[i] = make([]bool, size)
	}
	for i := 0; i < size-1; i++ {
		dp[i][i] = true
		dp[i][i+1] = str[i] == str[i+1]
	}
	dp[size-1][size-1] = true
	for j := 2; j < size; j++ {
		left := 0
		right := j
		for left < size && right < size {
			dp[left][right] = str[left] == str[right] && dp[left+1][right-1]
			left++
			right++
		}
	}
	return dp
}

func process(path []string, index int, str string, dp [][]bool) {
	if index == len(str) { // 收集答案
		cur := make([]string, len(path))
		copy(cur, path)
		result = append(result, cur)
	} else {
		for end := index; end < len(str); end++ {
			if dp[index][end] {
				// 标记现场
				path = append(path, str[index:end+1])
				// 进入下层决策
				process(path, end+1, str, dp)
				// 恢复现场
				path = path[:len(path)-1]
			}
		}
	}
}
