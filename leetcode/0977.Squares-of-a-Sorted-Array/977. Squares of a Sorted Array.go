package leetcode

import "sort"

// 解法一
// 双指针方式
// 1、由于存在负数，所以两头绝对值一定是最大的
// 2、通过双指针方式，将两头值比较，大的存入数组尾部，小的进行下一次比较
func sortedSquares2(A []int) []int {
	ans := make([]int, len(A))
	for i, k, j := 0, len(A)-1, len(ans)-1; i <= j; k-- {
		if A[i]*A[i] > A[j]*A[j] {
			ans[k] = A[i] * A[i]
			i++
		} else {
			ans[k] = A[j] * A[j]
			j--
		}
	}
	return ans
}

// 解法二
// 通过将值存放到当前值中，然后直接排序
func sortedSquares1(A []int) []int {
	for i, value := range A {
		A[i] = value * value
	}
	sort.Ints(A)
	return A
}
