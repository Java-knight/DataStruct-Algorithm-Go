package _123_best_time_to_buy_and_sell_stock_iii

// 买卖股票的最佳时机III
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/

// question: 整个过程只能买卖2次(但是每次手里只能持有一股 ——> 买了等卖了才能第二次买)
// 思路: 假设判断在i位置的值卖出, 需要在0~i-1上买卖一次的最优解减去最小买入股票的价格(把这个看成是一步操作)。
//      准备一次result收集每一步卖出的最优解
// ans = max(max(arr[0~i-1]) - arr[x]) + arr[i]; 这个公式就是需要得到i位置卖出完成两次买卖的最优解
// max(arr[0~i-1]): 这个就是在 0~i-1 位置上买卖一次的最优解;
// arr[x]: 0~i位置之间的一个最小值买入(第一次买入)
// max(max(arr[0~i-1]) - arr[x]): 最有还需要求出经过一次买卖, 再次入买入一笔的最大值
func maxProfit(prices []int) int {
	if prices == nil || len(prices) < 2 {
		return 0
	}
	// 解释: 一开始min是arr[0], 表示第一次买入的价格; 但是这里是在第一次买入直接卖出, 第二次买入也是arr[0]位置
	result := 0
	onceBuySell := 0                // max(arr[0~i-1])
	min := prices[0]                // arr[x], 在0号位置买入
	onceBuySellAndBuy := -prices[0] // max(max(arr[0~i-1]) - arr[x])
	for i := 1; i < len(prices); i++ {
		result = Max(result, onceBuySellAndBuy+prices[i])
		min = Min(min, prices[i])
		onceBuySell = Max(onceBuySell, prices[i]-min)
		onceBuySellAndBuy = Max(onceBuySellAndBuy, onceBuySell-prices[i])
	}
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
