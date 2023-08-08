package leetcode

func trap(height []int) int {
	left := 0                // 左指针初始位置
	right := len(height) - 1 // 右指针初始位置
	leftMax := 0             // 左边最大高度初始值
	rightMax := 0            // 右边最大高度初始值
	result := 0              // 总的接水量

	for left <= right { // 当左右指针没有交错时循环遍历
		if height[left] <= height[right] { // 左边高度小于等于右边高度，计算左边的接水量
			if height[left] >= leftMax {
				leftMax = height[left] // 更新左边最大高度
			} else {
				result += leftMax - height[left] // 计算接水量并累加到结果中
			}
			left++ // 左指针右移
		} else { // 右边高度小于左边高度，计算右边的接水量
			if height[right] >= rightMax {
				rightMax = height[right] // 更新右边最大高度
			} else {
				result += rightMax - height[right] // 计算接水量并累加到结果中
			}
			right-- // 右指针左移
		}
	}

	return result
}
