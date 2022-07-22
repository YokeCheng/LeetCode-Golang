package leetcode

// 方法一：循环替换
func reverseLeftWords(s string, n int) string {
	var res string
	s1 := []rune(s)

	res = string(append(s1[n:], s1[:n]...))

	return res
}
