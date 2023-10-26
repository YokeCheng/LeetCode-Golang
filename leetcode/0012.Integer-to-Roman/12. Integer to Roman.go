package leetcode

func intToRoman(num int) string {
	// 定义罗马数字的整数值和对应的符号
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := "" // 初始化结果字符串
	i := 0    // 初始化索引

	// 当数字不为0，循环操作
	for num != 0 {
		// 找到最大的整数值并添加对应的符号
		for values[i] > num {
			i++
		}
		num -= values[i]
		res += symbols[i]
	}

	return res // 返回结果字符串
}
