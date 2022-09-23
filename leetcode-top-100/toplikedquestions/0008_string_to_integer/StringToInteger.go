package main

import (
	"math"
	"strings"
)

// 字符串转换整数(atoi)
// https://leetcode.cn/problems/string-to-integer-atoi/

func myAtoi(s string) int {
	if s == "" {
		return 0
	}
	s = removeHeadZero(strings.TrimSpace(s))
	if s == "" {
		return 0
	}
	charArr := []rune(s)
	if !isValid(charArr) {
		return 0
	}
	// 是否是正数
	flag := ternaryBool(charArr[0] == '-', false, true)
	result := 0
	cur := 0
	minq := math.MinInt32 / 10
	minr := math.MinInt32 % 10
	for i := ternaryInt(charArr[0] == '-' || charArr[0] == '+', 1, 0); i < len(charArr); i++ {
		cur = int('0' - charArr[i])
		if result < minq || (result == minq && cur < minr) {
			return ternaryInt(flag, math.MaxInt32, math.MinInt32)
		}
		result = result*10 + cur
	}
	// 打补丁 case: "Integer.MAX_VALUE + 1" —> 会越界, 但是因为使用负数不会越界
	if flag && result == math.MinInt32 {
		return math.MaxInt32
	}
	return ternaryInt(flag, -result, result)
}

// 删除 str 头部的 ‘0’ 字符
func removeHeadZero(str string) string {
	flag := strings.HasPrefix(str, "+") || strings.HasPrefix(str, "-")
	index := ternaryInt(flag, 1, 0)
	charArr := []rune(str)
	for ; index < len(charArr); index++ {
		if charArr[index] != '0' {
			break
		}
	}

	end := -1
	for i := len(charArr) - 1; i >= index; i-- {
		if charArr[i] < '0' || charArr[i] > '9' {
			end = i
		}
	}
	end = ternaryInt(end == -1, len(charArr), end)
	var strHead string
	if flag {
		strHead = string(charArr[0])
	} else {
		strHead = ""
	}
	return strHead + str[index:end]
}

// 过滤器
func isValid(arr []rune) bool {
	if arr[0] != '-' && arr[0] != '+' && (arr[0] < '0' || arr[0] > '9') {
		return false
	}
	if (arr[0] == '-' || arr[0] == '+') && len(arr) == 1 {
		return false
	}
	for i := 1; i < len(arr); i++ {
		if arr[i] < '0' || arr[i] > '9' {
			return false
		}
	}
	return true
}

func ternaryBool(flag bool, a, b bool) bool {
	if flag {
		return a
	}
	return b
}

func ternaryInt(flag bool, a, b int) int {
	if flag {
		return a
	}
	return b
}
