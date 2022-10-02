package _122_best_time_to_buy_and_sell_stock_ii

// 买卖股票的最佳时机II
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/

// question: 整个过程能买卖无数次(但是每次手里只能持有一股 ——> 买了等卖了才能第二次买)
// 思路: 在i位置, 用arr[i] - arr[i-1] > 0, 表示需要买入
func maxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		ans := prices[i] - prices[i-1]
		if ans > 0 { // 收集答案
			result += ans
		}
	}
	return result
}
