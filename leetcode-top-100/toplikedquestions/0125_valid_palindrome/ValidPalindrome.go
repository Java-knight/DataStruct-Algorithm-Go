package _125_valid_palindrome

// 验证回文串
// https://leetcode.cn/problems/valid-palindrome/

// 思路: 直接搞左右两个指针, left向右, right向左;
// if s[left] != s[right] false, 否则left++, right--继续
func isPalindrome(s string) bool {
	if s == "" || len(s) == 0 {
		return true
	}

	left := 0
	right := len(s) - 1
	for left < right {
		if validChar(s[left]) && validChar(s[right]) { // 必须是大小写字母或数字
			if !equal(s[left], s[right]) { // 判断是否一样(字母/数字)
				return false
			}
			left++
			right--
		} else { // 符号/空格
			left += Ternary(validChar(s[left]), 0, 1)
			right -= Ternary(validChar(s[right]), 0, 1)
		}
	}
	return true
}

// 判断c是否是大小写字母或数字
func validChar(c byte) bool {
	return isLetter(c) || isNumber(c)
}

// 判断是否是大小写字母
func isLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

// 判断c是否是0~9之间的数字
func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

// 判断c1和c2是否一样(ascii表 A=65, a=97, 两者差32)
func equal(c1, c2 byte) bool {
	if isNumber(c1) || isNumber(c2) {
		return c1 == c2
	}
	return (c1 == c2) || (Max(c1, c2)-Min(c1, c2)) == 32
}

func Max(a, b byte) byte {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b byte) byte {
	if a > b {
		return b
	} else {
		return a
	}
}

func Ternary(flag bool, a, b int) int {
	if flag {
		return a
	} else {
		return b
	}
}
