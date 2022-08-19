package leetcode

func dicesProbability(n int) []float64 {
	dp := make([][]float64, n)
	for i := range dp {
		dp[i] = make([]float64, (i+1)*6-i)
	}
	for i := range dp[0] {
		dp[0][i] = float64(1) / float64(6)
	}
	for i := 1; i < len(dp); i++ {
		for j := range dp[i-1] {
			for k := range dp[0] {
				dp[i][j+k] += float64(dp[i-1][j]) * float64(dp[0][k])
			}
		}
	}
	return dp[n-1]
}
