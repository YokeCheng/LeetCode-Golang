package leetcode

import (
	"math"
	"strconv"
)

func printNumbers(n int) []int {
	var (
		buf, nums = make([]byte, n), make([]byte, 10)
		ret       = make([]int, 0, int(math.Pow10(n)))
		dfs       func(x, end int)
	)
	for i := 0; i < 10; i++ {
		nums[i] = '0' + byte(i)
	}
	dfs = func(x, end int) {
		if x == end {
			num, _ := strconv.Atoi(string(buf[:end]))
			ret = append(ret, num)
			return
		}
		var i int
		for ; i < 10; i++ {
			buf[x] = nums[i]
			dfs(x+1, end)
		}
	}
	for i := 1; i <= n; i++ { // 遍历所有长度
		dfs(0, i)
	}
	return ret[1:]
}
