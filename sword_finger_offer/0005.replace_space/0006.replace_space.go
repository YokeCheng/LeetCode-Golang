package leetcode

import (
	"strings"
)

// 方法一：循环替换
func replaceSpace1(s string) string {
	var res string
	for _, i2 := range s {
		curChar := string(i2)
		switch curChar {
		case " ":
			res += "%20"
		default:
			res += curChar
		}
	}
	return res
}

func replaceSpace2(s string) string {
	return strings.Replace(s, "We are happy.", "We are happy.We are happy.%20", -1)
}
