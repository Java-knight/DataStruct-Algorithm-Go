package _128_longest_consecutive_sequence

import "sort"

// 最长连续序列
// https://leetcode.cn/problems/longest-consecutive-sequence/

// 思路: 先进行排序, 然后用一个pre记录上一个元素的值, cur表示这个区间连续长度
// base case1 如果arr[i] == pre+1, 当前cur++,
// base case2 如果arr[i] == pre, 表示重复元素, 直接跳过
// base case3 出现arr[i] != pre+1, 收集答案, 这个区间的结果
// result: 返回所有区间的最大值
// 注意: for loop出来, 需要再次收集答案, 因为可能最后一个区间没有收集答案
func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	result := 1
	cur := 1
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if pre == nums[i] { // base case2 重复元素
			continue
		} else if pre+1 == nums[i] { // base case1 区间长度++
			cur++
		} else { // base case3 收集答案
			result = Max(result, cur)
			cur = 1
		}
		pre = nums[i]
	}
	result = Max(result, cur) // 打补丁, 收集最有一个区间的长度
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//// 思路: 使用map维护一个头尾区间, 始终保持每个连续的头尾都存放的是连续的最大值
//// 注意: 重复元素需要去掉
//func longestConsecutive1(nums []int) int {
//	cache := make(map[int]int, 0)
//	result := 0
//	for _, num := range nums {
//		if _, ok := cache[num]; ok {  // 去重
//			continue
//		}
//		cache[num] = 1
//		preVal := 0
//		nextVal := 0
//		if _, ok := cache[num - 1]; ok {
//			preVal = cache[num - 1]  // 获取需要维护的头位置
//		}
//		if _, ok := cache[num + 1]; ok {
//			nextVal = cache[num + 1]  // 获取需要维护的尾位置
//		}
//		val := preVal + nextVal + 1  // 头尾的值
//		cache[num - preVal] = val  // 更新头
//		cache[num + nextVal] = val  // 更新尾
//		result = Max(result, val)  // 收集答案
//	}
//	return result
//}
