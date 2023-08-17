package leetcode

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	num, lastInt, total := 0, 0, 0 // 初始化变量

	for i := len(s) - 1; i >= 0; i-- { // 反向遍历字符串
		char := s[i : i+1] // 获取当前字符
		num = roman[char]  // 查找字符对应的整数值

		if num < lastInt { // 当前数字小于上一个数字，需要进行减法运算
			total -= num
		} else { // 当前数字大于等于上一个数字，进行加法运算
			total += num
		}

		lastInt = num // 更新上一个数字
	}

	return total
}
