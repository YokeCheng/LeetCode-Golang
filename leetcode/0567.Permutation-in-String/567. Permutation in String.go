package leetcode

// 利用(字符串-'a')来制造map, 通过设置s1的字符map，当s2存在时则count--
// 利用双指针来保证字符串两者之间不会超出越界，即s2中判定的长度不会超过s1
func checkInclusion(s1 string, s2 string) bool {
	var freq [27]int
	if len(s2) == 0 || len(s2) < len(s1) {
		return false
	}
	// 形成s1的map，当存在之中则count--
	for i := 0; i < len(s1); i++ {
		freq[s1[i]-'a']++
	}
	left, right, count := 0, 0, len(s1)
	for right < len(s2) {
		if freq[s2[right]-'a'] >= 1 {
			count--
		}
		freq[s2[right]-'a']--
		right++
		if count == 0 {
			return true
		}
		if right-left == len(s1) {
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}
	return false
}
