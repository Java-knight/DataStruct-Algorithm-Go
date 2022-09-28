package _049_group_anagrams

import (
	"sort"
)

// 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/

// 同情字符串: 两个字符串中字母的数量和种类都相同就称为同情字符串, 比如str1="abc", str2="cba", str1和str2就是同情字符串
// 思路: 用一个map<String, List<String>>, key放的是经过排序后的str, value放的是同情字符串列表, 使用个list收集map中的values
func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	cache := make(map[string][]string)
	for _, str := range strs {
		charArr := []byte(str)
		sort.Slice(charArr, func(i, j int) bool { // 排序(耗时大)
			return charArr[i] < charArr[j]
		})
		key := string(charArr)
		if _, ok := cache[key]; !ok {
			cache[key] = make([]string, 0)
		}
		cache[key] = append(cache[key], str)
	}
	for _, val := range cache {
		result = append(result, val)
	}
	return result
}
