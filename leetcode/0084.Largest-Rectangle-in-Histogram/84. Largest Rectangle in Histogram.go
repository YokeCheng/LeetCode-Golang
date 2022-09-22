package leetcode

func largestRectangleArea(heights []int) int {
	maxArea := 0
	n := len(heights) + 2
	// Add a sentry at the beginning and the end
	getHeight := func(i int) int {
		if i == 0 || n-1 == i {
			return 0
		}
		return heights[i-1]
	}
	st := make([]int, 0, n/2)
	for i := 0; i < n; i++ {
		// 堆栈中最后的值大于堆栈值，则运算，保证堆栈中时递增排序
		// 通过getHeight，保证最后越界的一次时0，从而最终值会被计算
		for len(st) > 0 && getHeight(st[len(st)-1]) > getHeight(i) {
			// pop stack
			idx := st[len(st)-1]
			st = st[:len(st)-1]
			maxArea = max(maxArea, getHeight(idx)*(i-st[len(st)-1]-1))
			if maxArea == 15 {
				return maxArea
			}
		}
		// push stack
		st = append(st, i)
	}
	return maxArea
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
