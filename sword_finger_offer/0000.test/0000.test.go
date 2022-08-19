package leetcode

func findContinuousSequence(target int) [][]int {
	var res [][]int
	for n := target; n > 0; n-- {
		sub := (n + 1) * n / 2
		if target-sub < 0 {
			continue
		}
		if (target-sub)%(n+1) == 0 {
			a := (target - sub) / (n + 1)
			var b []int
			for i := 0; i <= n; i++ {
				b = append(b, a+i)
			}
			res = append(res, b)
		}
	}
	return res
}
