package leetcode

func maxValue(grid [][]int) int {
	row := len(grid) - 1
	col := len(grid[0]) - 1
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	return max(left(row, col, grid, dp), top(row, col, grid, dp)) + grid[row][col]
}

func left(row, col int, grid, dp [][]int) int {
	curRow := row - 1
	if curRow < 0 {
		return 0
	}
	if dp[curRow][col] <= 0 {
		dp[curRow][col] = max(left(curRow, col, grid, dp), top(curRow, col, grid, dp)) + grid[curRow][col]
	}
	return dp[curRow][col]
}

func top(row, col int, grid, dp [][]int) int {
	curCol := col - 1
	if curCol < 0 {
		return 0
	}
	if dp[row][curCol] <= 0 {
		dp[row][curCol] = max(left(row, curCol, grid, dp), top(row, curCol, grid, dp)) + grid[row][curCol]
	}
	return dp[row][curCol]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
