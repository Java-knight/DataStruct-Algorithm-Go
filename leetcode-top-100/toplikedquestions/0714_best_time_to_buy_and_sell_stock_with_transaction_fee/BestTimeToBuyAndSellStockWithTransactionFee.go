package _714_best_time_to_buy_and_sell_stock_with_transaction_fee

import "math"

// 买卖股票的最佳时机含有手续费
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

// question: 整个过程能买卖无数次, 但是在买入时候需要支付一次手续费, 卖出是不需要的
// 思路: 在i位置交易的时候, 需要提前保存一下i-1位置的卖出价格, 用于在i位置买入做计算
// dp[i]sell = max(dp[i-1]sell, dp[i-1]buy+arr[i])
// dp[i]buy = max(dp[i-1]buy, dp[i-1]sell-arr[i]-fee)
func maxProfit(prices []int, fee int) int {
	sell := 0
	buy := math.MinInt32
	for i := 0; i < len(prices); i++ {
		tmp := sell                       // 因为买入会产生手续费, 所以需要提前保存下卖出的价格
		sell = Max(sell, buy+prices[i])   // 卖出
		buy = Max(buy, tmp-prices[i]-fee) // 买入(手续费)
	}
	return sell
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
