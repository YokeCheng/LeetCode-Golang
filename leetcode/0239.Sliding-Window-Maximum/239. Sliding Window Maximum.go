package leetcode

// 解法一 暴力解法 O(nk)
func maxSlidingWindow1(a []int, k int) []int {
	res := make([]int, 0, k)
	n := len(a)
	if n == 0 {
		return []int{}
	}
	for i := 0; i <= n-k; i++ {
		max := a[i]
		for j := 1; j < k; j++ {
			if max < a[i+j] {
				max = a[i+j]
			}
		}
		res = append(res, max)
	}

	return res
}

// 解法二 双端队列 Deque
func maxSlidingWindow(nums []int, k int) []int {
	var queue []int
	var res []int
	for i := 0; i < len(nums); i++ {
		// 循环将最后一个值小于当前值的剔除，其实相当于保留了一个最大值在最前面
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
		// 通过队列中第一个数和 前第K个数比较,相同表示本轮窗格应该去掉开始的数
		if i >= k && queue[0] == nums[i-k] {
			queue = queue[1:]
		}
		// 由于是递增的，所以直接取最开始的数即可
		if i >= k-1 {
			res = append(res, queue[0])
		}
	}
	return res
}
