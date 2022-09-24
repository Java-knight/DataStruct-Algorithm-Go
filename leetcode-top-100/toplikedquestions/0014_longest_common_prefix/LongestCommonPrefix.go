package _014_longest_common_prefix

import "math"

// 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	charArr := strs[0]
	min := math.MaxInt32
	for _, str := range strs {
		index := 0
		for index < len(str) && index < len(charArr) {
			if str[index] != charArr[index] {
				break
			}
			index++
		}
		min = Min(min, index)
	}
	return charArr[0:min]
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
