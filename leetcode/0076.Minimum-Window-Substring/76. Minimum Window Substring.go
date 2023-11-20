package leetcode

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	// 用于记录字符频率的数组，tFreq 用于存储 t 中每个字符的频率，sFreq 用于存储当前窗口中每个字符的频率
	var tFreq, sFreq [58]int

	// 用于存储最终结果的变量
	var result string
	// 窗口的左右边界，finalLeft 和 finalRight 分别表示最小窗口子串的左右边界
	left, right, finalLeft, finalRight := 0, 0, -1, -1

	// 记录窗口中包含 t 中字符的个数
	count := 0

	// 初始化 t 中每个字符的频率
	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'A']++
	}

	// 滑动窗口算法
	for right < len(s) {
		// 更新当前窗口中右边字符的频率
		sFreq[s[right]-'A']++

		// 如果该字符的频率不超过 t 中对应字符的频率，则 count 增加
		if sFreq[s[right]-'A'] <= tFreq[s[right]-'A'] {
			count++
		}

		// 如果当前窗口包含 t 中所有字符
		for count == len(t) {
			// 更新最小窗口的左右边界
			if finalLeft == -1 || right-left < finalRight-finalLeft {
				finalLeft, finalRight = left, right
			}

			// 移动左边界
			// 如果移动左边界后，移出的字符是 t 中的字符，则 count 减少
			sFreq[s[left]-'A']--
			if sFreq[s[left]-'A'] < tFreq[s[left]-'A'] {
				count--
			}
			left++
		}

		// 移动右边界
		right++
	}

	// 根据最终左右边界提取最小窗口子串
	if finalLeft != -1 {
		result = string(s[finalLeft : finalRight+1])
	}
	return result
}

func minWindow2(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	tFreq, sFreq := make(map[byte]int), make(map[byte]int)

	// 初始化 t 中每个字符的频率
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	left, right, finalLeft, finalRight := 0, 0, -1, -1
	count := 0

	for right < len(s) {
		// 更新当前窗口中右边字符的频率
		sFreq[s[right]]++

		// 如果该字符的频率不超过 t 中对应字符的频率，则 count 增加
		if sFreq[s[right]] <= tFreq[s[right]] {
			count++
		}

		// 如果当前窗口包含 t 中所有字符
		for count == len(t) {
			// 更新最小窗口的左右边界
			if finalLeft == -1 || right-left < finalRight-finalLeft {
				finalLeft, finalRight = left, right
			}

			// 移动左边界
			// 如果移动左边界后，移出的字符是 t 中的字符，则 count 减少
			sFreq[s[left]]--
			if sFreq[s[left]] < tFreq[s[left]] {
				count--
			}
			left++
		}

		// 移动右边界
		right++
	}

	// 根据最终左右边界提取最小窗口子串
	if finalLeft != -1 {
		return s[finalLeft : finalRight+1]
	}
	return ""
}
