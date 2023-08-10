package leetcode

func jump2(nums []int) int {
	n := len(nums) - 1
	if n == 0 {
		return 0
	}

	steps := 0       // 跳跃次数
	maxPosition := 0 //最远位置
	end := 0         // 当前边界

	for i, num := range nums {
		maxPosition = max(i+num, maxPosition)
		if maxPosition >= n {
			return steps
		}
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func jump(nums []int) int {
	n := len(nums) - 1

	// 定义变量
	steps := 0       // 跳跃次数
	maxPosition := 0 // 当前能够跳跃的最远位置
	end := 0         // 当前能够跳跃的边界

	// 假设了一定能跳到最后，则设置了小于当前值来保证最后一步不用跳，因为end必定大于i
	for i := 0; i < n; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if maxPosition >= n {
			return steps + 1
		}
		// 如果到达边界，更新边界值，并增加跳跃次数
		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
