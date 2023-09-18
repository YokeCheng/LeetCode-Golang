package leetcode

var roman = map[byte]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
}

func romanToInt(s string) int {
	total := 0 // 初始化总数

	prev := 0 // 初始化前一个字符的值为0

	for i := len(s) - 1; i >= 0; i-- { // 反向遍历字符串
		current := roman[s[i]] // 获取当前字符对应的整数值

		if current >= prev { // 如果当前数字大于前一个数字
			total += current // 需要加上前一个数字
		} else {
			total -= current // 否则减去当前数字
		}
		prev = current
	}

	return total
}
