package _056_merge_intervals

import "sort"

// 合并区间
// https://leetcode.cn/problems/merge-intervals/

// 思路: intervals比是一个Nx2的数组, 先按照 Range.start进行排序(比较器), 再根据两个base case 进行合并
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return make([][]int, 0)
	}
	arr := make([]Range, len(intervals))
	for i, val := range intervals {
		arr[i] = Range{val[0], val[1]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].start < arr[j].start
	})
	result := make([][]int, 0)
	s := arr[0].start
	e := arr[0].end
	for i := 1; i < len(intervals); i++ {
		if arr[i].start > e {
			result = append(result, []int{s, e}) // 收集答案
			s = arr[i].start
			e = arr[i].end
		} else { // 更新范围
			e = Max(e, arr[i].end)
		}
	}
	result = append(result, []int{s, e}) // 收集最后一个
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Range struct {
	start int
	end   int
}
