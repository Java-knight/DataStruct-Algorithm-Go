package main

// 最长无重复子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/?favorite=2cktkvj

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	charArr := []rune(s)
	cache := make([]int, 256)
	for i := 0; i < 256; i++ {
		cache[i] = -1
	}
	res := 0
	pre := -1
	cur := 0
	for i := 0; i < len(charArr); i++ {
		pre = max(pre, cache[charArr[i]])
		cur = i - pre
		res = max(res, cur)
		cache[charArr[i]] = i
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
