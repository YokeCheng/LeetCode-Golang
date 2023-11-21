package leetcode

// 利用(字符串-'a')来制造map, 通过设置s1的字符map，当s2存在时则count--
// 利用双指针来保证字符串两者之间不会超出越界，即s2中判定的长度不会超过s1

// 获取短字符串S的长度count
// 通过左右指针来维护一个和字符串S相同长度的窗格
// 通过长的字符串L去比较短的字符串S，比较后存在值，count-1
// 超过长度就进行窗格右移动，但是移动时需要注意，剔除窗格的值存在S中时要对count进行处理

func checkInclusion(s1 string, s2 string) bool {
	// freq 数组用于统计 s1 中每个字符的出现次数
	var freq [27]int

	// 如果 s2 的长度为 0 或小于 s1 的长度，直接返回 false
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}

	// 初始化 freq 数组，统计 s1 中每个字符的出现次数
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}

	// 使用双指针滑动窗口进行匹配
	left, right, count := 0, 0, len(s1)

	for right < len(s2) {
		// 如果当前字符在 s1 中存在，则 count 减一
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		// 更新频率数组
		freq[s2[right]-'a']--
		right++

		// 如果 count 减到 0，说明找到了一个符合条件的子串
		if count == 0 {
			return true
		}

		// 如果窗口大小等于 s1 的长度，需要移动左指针
		if right-left == len(s1) {
			// 如果左指针所指的字符在 s1 中存在，则 count 加一
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			// 恢复频率数组
			freq[s2[left]-'a']++
			left++
		}
	}

	// 如果没有找到符合条件的子串，返回 false
	return false
}
