package leetcode

// 假设成密码锁来看，
func countDigitOne(n int) int {
	// cur 为当前位的值
	// base 当前位数
	var base, res int64 = 1, 0
	n64 := int64(n)
	for base <= n64 {
		// 计算 a值，a值代表当前值前面的值
		a := n64 / base / 10
		// 取当前值
		cur := (n64 / base) % 10
		// 计算b值，b值代表当前值后面的值
		b := n64 % base
		// 将当前位设为1，考察其他部分的变化范围
		if cur > 1 {
			res += (a + 1) * base
		} else if cur == 1 {
			res += a*base + b + 1
		} else {
			res += a * base
		}
		// 统计更高一位
		base *= 10
	}
	return int(res)
}
