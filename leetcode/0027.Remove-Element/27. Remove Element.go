package leetcode

import "sort"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	// j永远获取的是最开始的等于val值的下标
	// i为当前值，不等于val时则和j调换位置
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}

// 排序后查找位置，再去拼接
func removeElement2(nums []int, val int) int {
	sort.Ints(nums)
	left, right := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			right++
		} else if left == right {
			left++
			right++
		} else {
			break
		}
	}
	nums = append(nums[:left], nums[right:]...)
	return len(nums)
}
