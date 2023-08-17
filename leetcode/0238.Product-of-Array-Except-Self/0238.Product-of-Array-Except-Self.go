package leetcode

// 解法一
// 时间复杂度O(N)   空间复杂度O(N)
// L 和 R 分别表示左右两侧的乘积列表,
// L计算了当前值左边的乘积， R计算了当前值右边的乘积
// 两个相乘则显示了除当前值以外其他值的乘积
func productExceptSelf1(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}

	// L 和 R 分别表示左右两侧的乘积列表
	L, R, answer := make([]int, length), make([]int, length), make([]int, length)

	// L[i] 为索引 i 左侧所有元素的乘积
	// 对于索引为 '0' 的元素，因为左侧没有元素，所以 L[0] = 1
	L[0] = 1
	for i := 1; i < length; i++ {
		L[i] = nums[i-1] * L[i-1]
	}

	// R[i] 为索引 i 右侧所有元素的乘积
	// 对于索引为 'length-1' 的元素，因为右侧没有元素，所以 R[length-1] = 1
	R[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	// 对于索引 i，除 nums[i] 之外其余各元素的乘积就是左侧所有元素的乘积乘以右侧所有元素的乘积
	for i := 0; i < length; i++ {
		answer[i] = L[i] * R[i]
	}
	return answer
}

// 解法二
// 时间复杂度O(N)   空间复杂度O(N) 当前空间比上一个方法小一倍
// answer表示左侧的乘积列表,
// answer计算了当前值左边的乘积数组， R计算了当前值右边的乘积
// 通过右侧循环直接计算值，这样就可以减少对空间的使用
func productExceptSelf(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)

	// answer[i] 表示索引 i 左侧所有元素的乘积
	// 因为索引为 '0' 的元素左侧没有元素， 所以 answer[0] = 1
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}

	// R 为右侧所有元素的乘积
	// 刚开始右边没有元素，所以 R = 1
	R := 1
	for i := length - 1; i >= 0; i-- {
		// 对于索引 i，左边的乘积为 answer[i]，右边的乘积为 R
		answer[i] = answer[i] * R
		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
		R *= nums[i]
	}
	return answer
}
