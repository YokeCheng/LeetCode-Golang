package leetcode

import (
	"math"
)

func coinChange(coins []int, amount int) int {
	// 创建一个动态规划数组 dp，长度为 amount+1，用于存储每个金额对应的最小硬币数量
	dp := make([]int, amount+1)
	// 初始化 dp[0] 为 0，因为组成金额 0 的最小硬币数量为 0
	dp[0] = 0

	// 初始化 dp 数组的其余元素，将它们都设置为 amount+1
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	// 外部循环遍历每个金额 i 从 1 到 amount
	for i := 1; i <= amount; i++ {
		// 内部循环遍历所有给定的硬币面额
		for j := 0; j < len(coins); j++ {
			// 检查硬币面额 coins[j] 是否小于或等于当前的金额 i
			if coins[j] <= i {
				// 尝试更新 dp[i]，通过选择 dp[i-coins[j]] 和当前 dp[i] 中的最小值，然后加上 1
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	// 如果 dp[amount] 仍然等于 amount+1，表示无法组成所需的金额，返回 -1
	if dp[amount] > amount {
		return -1
	}
	// 否则，返回 dp[amount]，它包含了用给定硬币组成所需金额的最小硬币数量
	return dp[amount]
}

// 辅助函数 min 用于返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
