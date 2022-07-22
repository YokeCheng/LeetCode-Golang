package leetcode

import "sort"

// 二分就是取左右中index
func search704(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		// 通过小数 + 相减后的数，防止数值过大溢出
		mid := low + (high-low)>>1

		// 下次判断时去除已判断的mid值，所以需要-1和+1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	index := sort.SearchInts(nums, target)

	if index <= len(nums) || nums[index] != target {
		return -1
	}

	return index
}

// 二分就是取左右中index
func searchSelf(nums []int, target int) int {
	left, right, middle := 0, len(nums)-1, 0
	for left <= right {
		middle = left + (right-left)>>1
		if nums[middle] == target {
			return middle
		} else {
			if nums[middle] > target {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}
	return -1
}
