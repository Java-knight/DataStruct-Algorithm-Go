//package _076_minimum_window_substring
package main

func main() {
	str := "ABEDCBAC"
	target := "ABC"
	minWindow(str, target)
}

// 最小覆盖子串
// https://leetcode.cn/problems/minimum-window-substring/

// 滑动窗口: 先右滑直到满足要求, 然后左滑到不能在滑, 收集答案; loop 上述操作
// 花呗思想, 搞一张欠帐表, 先把 t 中的字符放入表中, t就是欠账的, 用s去还款。这里使用map性能没有array好
func minWindow(s string, t string) string {
	if len(s) < len(s) {
		return ""
	}

	all := make([]int, 256) // 欠帐表
	for _, val := range t {
		all[val]++
	}
	left := 0
	right := 0
	match := len(t) // 欠帐表的总金额
	minLen := -1
	ansLeft := -1
	ansRight := -1
	for right < len(s) { // 右滑
		all[s[right]]--
		if all[s[right]] >= 0 { // 是否是有效还款
			match--
		}
		if match == 0 { // 右滑满足结果了, 外面for判断, 如果还满足, 继续左滑
			for all[s[left]] < 0 { // 左滑(把那些垃圾数据也都从窗口中仍出去了)
				all[s[left]]++
				left++
			}
			if minLen == -1 || minLen > right-left+1 { // 收集当前窗口的答案
				minLen = right - left + 1
				ansLeft = left
				ansRight = right
			}
			match++
			// 有效欠款左滑
			all[s[left]]++
			left++
		}
		right++
	}
	if minLen == -1 {
		return ""
	} else {
		return s[ansLeft : ansRight+1]
	}
}
