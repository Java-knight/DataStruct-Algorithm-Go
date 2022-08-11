package leetcode

// https://leetcode.cn/problems/two-sum/
/* 注解: cache<val, key>, 需要返回的是index */
func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for key, val := range nums {
		if index, ok := cache[target-val]; ok {
			return []int{index, key}
		}
		cache[val] = key
	}
	return nil
}
