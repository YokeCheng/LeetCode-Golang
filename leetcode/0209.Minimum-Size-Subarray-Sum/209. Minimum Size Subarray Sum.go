package leetcode

// 巧用for循环来取代递归操作
func minSubArrayLen(target int, nums []int) int {
	left, sum, res := 0, 0, len(nums)+1
	for right, v := range nums {
		sum += v
		// 巧用for循环来取代递归操作
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
