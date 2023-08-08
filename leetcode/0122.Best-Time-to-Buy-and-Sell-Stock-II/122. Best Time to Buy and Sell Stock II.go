package leetcode

// 买卖股票最佳时机
// 通过判断后面是否大于前面的值，只要大于，则适合买入
// 由于不需要确定买入卖出的时间，则可以直接理解为，当后一天大于前一天，则前一天买入，第二天卖出即可
func maxProfit122(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}
