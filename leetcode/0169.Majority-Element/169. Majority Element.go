package leetcode

import "sort"

// 解法一 时间复杂度 O(n) 空间复杂度 O(1)
func majorityElement(nums []int) int {
	res, count := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res, count = nums[i], 1
		} else {
			if nums[i] == res {
				count++
			} else {
				count--
			}
		}
	}
	return res
}

// 解法二 时间复杂度 O(n) 空间复杂度 O(n)
func majorityElement1(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

// 解法三 排序取中法
func majorityElement3(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	return nums[n/2]
}

// https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof/solution/gou-jian-cheng-ji-shu-zu-by-leetcode-sol-aqg2/
//
func constructArr(a []int) []int {
	length := len(a)

	// L 和 R 分别表示左右两侧的乘积列表
	L, R, answer := make([]int, length), make([]int, length), make([]int, length)

	// L[i] 为索引 i 左侧所有元素的乘积
	// 对于索引为 '0' 的元素，因为左侧没有元素，所以 L[0] = 1
	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = a[i-1] * L[i-1]
	}

	// R[i] 为索引 i 右侧所有元素的乘积
	// 对于索引为 'length-1' 的元素，因为右侧没有元素，所以 R[length-1] = 1
	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = a[i+1] * R[i+1]
	}

	// 对于索引 i，除 a[i] 之外其余各元素的乘积就是左侧所有元素的乘积乘以右侧所有元素的乘积
	for i := 0; i < length; i++ {
		answer[i] = L[i] * R[i]
	}
	return answer
}
