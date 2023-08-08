package leetcode

// 解法一 模拟 DP(动态规划)
// 找到 当前值-最小值 > max的值
// 比较当前值 和 最小值min的大小，保证永远后面的值都是和最小值比较
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, maxProfitVal := prices[0], 0
	for _, price := range prices {
		if maxProfitVal < price-min {
			maxProfitVal = price - min
		}
		if min < price {
			min = price
		}
	}
	return maxProfitVal
}

// 解法二 单调栈
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	stack, res := []int{prices[0]}, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > stack[len(stack)-1] {
			stack = append(stack, prices[i])
		} else {
			index := len(stack) - 1
			for ; index >= 0; index-- {
				if stack[index] < prices[i] {
					break
				}
			}
			stack = stack[:index+1]
			stack = append(stack, prices[i])
		}
		res = max(res, stack[len(stack)-1]-stack[0])
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
