package leetcode

func lengthOfLastWord(s string) int {
	last := len(s) - 1

	// 从字符串末尾开始，跳过末尾的空格
	for last >= 0 && s[last] == ' ' {
		last--
	}

	// 如果字符串全为空格，则返回长度为 0
	if last < 0 {
		return 0
	}

	first := last

	// 从末尾的空格处向前查找，找到最后一个单词的起始位置
	for first >= 0 && s[first] != ' ' {
		first--
	}

	// 返回最后一个单词的长度
	return last - first
}
