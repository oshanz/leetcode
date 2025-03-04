func maxProfit(prices []int) int {
	var max_profit = 0
	var l = len(prices)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l && (prices[j] > prices[i]); j++ {
			var profit = prices[j] - prices[i]
			if profit > max_profit {
				max_profit = profit
			}
		}
	}
	return max_profit
}
