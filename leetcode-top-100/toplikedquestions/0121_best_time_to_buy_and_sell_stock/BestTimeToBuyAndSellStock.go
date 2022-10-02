package _121_best_time_to_buy_and_sell_stock

import "math"

// 买卖股票的最佳时机
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

// question: 整个过程只能买卖一次
// 思路: 假设第i天卖出, 那么就是0~i天最少价格的一天买入min(arr[0~i]);
// 最终结果就是max(arr[0~len-1])卖出的结果
func maxProfit(prices []int) int {
	max := math.MinInt32
	min := math.MaxInt32
	for _, val := range prices {
		min = Min(min, val) // 获取arr[0~i]上的最小值
		ans := val - min    // 保存第i天卖出的最优结果
		max = Max(max, ans) // 收集最终
	}
	return max
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
