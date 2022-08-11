package leetcode

// 递归 时间复杂度 O(log n),空间复杂度 O(1)
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	tmp := myPow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	}
	return tmp * tmp * x
}

// 递归b
func myPow2(x float64, b int) float64 {
	if x == 0 {
		return 0
	}
	var res = 1.0
	if b < 0 {
		x = 1 / x
		b = -b
	}
	for b > 0 {
		// 判断是否是奇数
		if (b & 1) == 1 {
			res *= x
		}
		x *= x
		b >>= 1
	}
	return res
}
