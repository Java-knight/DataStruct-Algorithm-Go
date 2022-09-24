//package _017_letter_combination_of_a_phone_number
package main

import "log"

// 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

var phone = [][]rune{
	{'a', 'b', 'c'},      // 2  0(按键, 下标)
	{'d', 'e', 'f'},      // 3  1
	{'g', 'h', 'i'},      // 4  2
	{'j', 'k', 'l'},      // 5  3
	{'m', 'n', 'o'},      // 6  4
	{'p', 'q', 'r', 's'}, // 7  5
	{'t', 'u', 'v'},      // 8  6
	{'w', 'x', 'y', 'z'}, // 9  7
}

var result []string // 定义在外面的原因, 因为需要让整个递归过程中都可以取到同一份result

func letterCombinations(digits string) []string {
	result = make([]string, 0)
	if digits == "" || len(digits) == 0 {
		return result
	}
	charArr := []rune(digits)
	process(charArr, 0, "")
	return result
}

func process(charArr []rune, index int, path string) {
	if index == len(charArr) { // 收集答案
		result = append(result, path)
	} else {
		cands := phone[charArr[index]-'2']
		for _, cur := range cands {
			// 保存现场
			tmp := path + string(cur)
			// 进入下层决策树
			process(charArr, index+1, tmp)
			// 恢复现场(不需要, 因为tmp是一个临时的)
		}
	}
}

func main() {
	log.Println(letterCombinations("23"))
}
