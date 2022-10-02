package _309_best_time_to_buy_and_sell_stock_with_cooldown

// 最佳买卖股票时机含冷冻期
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/

// question: 整个过程能买卖无数次, 但是在卖出后有一天的冷冻期(cooldown)
// 思路: 保存一下, 第i天两天前卖出sellPre、买入buy、卖出sell他们的最优解, 最后返回sell
func maxProfit(prices []int) int {
	size := len(prices)
	sellPre := 0      // 卖出股票(两天前, 已经可以进行买入了)
	buy := -prices[0] // 买入股票
	sell := 0         // 卖出股票
	for i := 1; i < size; i++ {
		tmp := sell
		// 这个买入/卖出 就是那一天的冷冻期
		sell = Max(sell, buy+prices[i])
		buy = Max(buy, sellPre-prices[i])
		sellPre = tmp
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
