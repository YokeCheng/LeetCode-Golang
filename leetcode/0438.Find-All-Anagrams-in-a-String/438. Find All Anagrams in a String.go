package leetcode

func findAnagrams(s string, p string) []int {
	var freq [26]int
	result := []int{}

	// 特殊情况处理
	if len(s) == 0 || len(s) < len(p) {
		return result
	}

	// 初始化频次数组
	for i := 0; i < len(p); i++ {
		freq[p[i]-'a']++
	}

	left, right, count := 0, 0, len(p)

	// 遍历 s 字符串
	for right < len(s) {
		// 右侧字符在 p 中存在，count 减一
		if freq[s[right]-'a'] >= 1 {
			count--
		}
		freq[s[right]-'a']--
		right++

		// 当前窗口中的字符都匹配完成
		if count == 0 {
			result = append(result, left)
		}

		// 窗口长度达到 len(p)，移动左侧
		if right-left == len(p) {
			// 左侧字符是 p 中的有效字符，count 加一
			if freq[s[left]-'a'] >= 0 {
				count++
			}
			freq[s[left]-'a']++
			left++
		}
	}

	return result
}
