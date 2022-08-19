package leetcode

func reversePairs(nums []int) int {
	res := 0
	temp := make([]int, len(nums))
	for l := 1; l < len(nums); l <<= 1 {
		for i := 0; i+l < len(nums); i += 2 * l {
			res += merge(nums, temp, i, i+l-1, min(len(nums)-1, i+2*l-1))
		}
	}
	return res
}

func merge(nums, temp []int, l, m, r int) int {
	l1, l2, r1, r2, k := l, m+1, m, r, l
	cnt := 0
	for l1 <= r1 && l2 <= r2 {
		if nums[l1] <= nums[l2] {
			temp[k] = nums[l1]
			l1++
			k++
		} else {
			cnt += r1 - l1 + 1
			temp[k] = nums[l2]
			l2++
			k++
		}
	}
	for l1 <= r1 {
		temp[k] = nums[l1]
		l1++
		k++
	}
	for l2 <= r2 {
		temp[k] = nums[l2]
		l2++
		k++
	}
	for i := l; i <= r; i++ {
		nums[i] = temp[i]
	}
	return cnt
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
