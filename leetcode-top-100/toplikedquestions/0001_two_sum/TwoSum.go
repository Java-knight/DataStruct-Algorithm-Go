package main

// 两数之和
// https://leetcode.cn/problems/two-sum/?favorite=2cktkvj

func twoSum(nums []int, target int) []int {
	// cache<val, index>
	cache := make(map[int]int, 0)
	for i, val := range nums {
		if index, ok := cache[target-val]; ok {
			return []int{index, i}
		}
		cache[val] = i
	}
	return []int{-1, -1}
}

func main() {

}
