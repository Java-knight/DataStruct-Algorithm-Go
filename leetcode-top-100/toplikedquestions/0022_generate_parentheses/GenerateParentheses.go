//package _022_generate_parentheses
package main

import "fmt"

// 括号生成 (如果只要求数量, 那就是科特兰数)
// https://leetcode.cn/problems/generate-parentheses/

func main() {
	fmt.Println(generateParenthesis(3))
}

var result []string

func generateParenthesis(n int) []string {
	result = make([]string, 0)
	path := make([]byte, n<<1)
	process(path, 0, 0, n)
	return result
}

/**
* 暴力递归
* 比如: n = 5, index = 4
*  path 当前路径  "()(("
*  index 当前路径的位置;
*  leftMatch 待匹配的左括号; 1, 当前路径有一个左括号待匹配
*  leftRest 剩余可以使用的左括号 2, 当前路径最多还能增加一个左括号
*  result 收集路径集合的结果
 */
func process(path []byte, index, leftMatch, leftRest int) {
	if index == len(path) {
		result = append(result, string(path))
	} else {
		if leftRest > 0 { // 填入左括号
			path[index] = '('
			process(path, index+1, leftMatch+1, leftRest-1)
		}
		if leftMatch > 0 { // 填入右括号
			path[index] = ')'
			process(path, index+1, leftMatch-1, leftRest)
		}
	}
}

// 暴力递归, 没有剪枝
//func generateParenthesis(n int) []string {
//	result = make([]string, 0)
//	path := make([]rune, n << 1)
//	process(path, 0)
//	return result
//}
//
//func process(path []rune, index int) {
//	if index == len(path) {
//		if isValid(path) {
//			result = append(result, string(path))
//		}
//	} else {
//		path[index] = '('
//		process(path, index + 1)
//		path[index] = ')'
//		process(path, index + 1)
//	}
//}
//
//func isValid(path []rune) bool {
//	count := 0
//	for _, val := range path{
//		if val == '(' {
//			count++
//		} else {
//			count--
//		}
//		if count < 0 {
//			return false
//		}
//	}
//	return count == 0
//}
