package leetcode

// lengthOfLongestSubstring6 寻找最长的无重复字符子串的长度。
func lengthOfLongestSubstring6(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	curStart, res := 0, 0
	indexes := make(map[rune]int, len(s))

	// 遍历字符串
	for i, a := range s {
		// 检查字符是否已经在子串中
		if index, ok := indexes[a]; ok {
			if index >= curStart {
				// 如果是，更新当前子串的起始位置
				curStart = indexes[a] + 1
			}
		}

		// 更新当前子串的长度
		if res < i-curStart+1 {
			res = i - curStart + 1
		}

		// 更新当前字符的索引
		indexes[a] = i
	}
	return res
}

// lengthOfLongestSubstring1 使用位图寻找最长的无重复字符子串的长度。
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var bitSet [26]bool
	result, left, right := 0, 0, 0

	// 遍历字符串
	for left < len(s) {
		// 检查右侧字符是否已经在子串中
		if s[right]-'a' < 26 && bitSet[s[right]-'a'] {
			// 如果是，将左指针向前移动，直到字符不在子串中
			if s[left]-'a' < 26 {
				bitSet[s[left]-'a'] = false
			}
			left++
		} else {
			if s[right]-'a' < 26 {
				// 如果不是，更新子串并将右指针向前移动
				bitSet[s[right]-'a'] = true
			}
			right++
		}

		// 更新当前子串的长度
		if result < right-left {
			result = right - left
		}

		// 如果左指针加上当前长度已经超过字符串长度，或者右指针已经达到字符串末尾，跳出循环
		if left+result >= len(s) || right >= len(s) {
			break
		}
	}
	return result
}

// lengthOfLongestSubstring5 使用滑动窗口寻找最长的无重复字符子串的长度。
func lengthOfLongestSubstring5(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1

	// 遍历字符串
	for left < len(s) {
		// 检查向前移动右指针是否保持无重复字符的子串
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			// 如果不是，将左指针向前移动
			freq[s[left]]--
			left++
		}

		// 更新当前子串的长度
		result = max(result, right-left+1)
	}
	return result
}

// lengthOfLongestSubstring2 使用滑动窗口和哈希桶寻找最长的无重复字符子串的长度。
func lengthOfLongestSubstring2(s string) int {
	right, left, res := 0, 0, 0
	indexes := make(map[byte]int, len(s))

	// 遍历字符串
	for left < len(s) {
		// 检查左侧字符是否已经在子串中
		if idx, ok := indexes[s[left]]; ok && idx >= right {
			// 如果是，更新当前子串的起始位置
			right = idx + 1
		}

		// 更新当前字符的索引
		indexes[s[left]] = left
		left++

		// 更新当前子串的长度
		res = max(res, left-right)
	}
	return res
}

// max 返回两个整数的最大值。
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
