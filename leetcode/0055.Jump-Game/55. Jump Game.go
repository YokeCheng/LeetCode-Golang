package leetcode

// maxJump是循环切片获取的之前跳跃能到达的最远位置
// 当某一个下标大于maxJump，则表示无法到达此位置，失败
func canJump2(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	maxJump := 0
	for i, v := range nums {
		if i > maxJump {
			return false
		}
		maxJump = max(maxJump, i+v)
		if maxJump > len(nums)-1 {
			return true
		}
	}
	return true
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
