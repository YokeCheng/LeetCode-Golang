package leetcode

func findMin154(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}
		mid := low + (high-low)>>1
		if nums[mid] > nums[low] {
			low = mid + 1
		} else if nums[mid] == nums[low] {
			// 相等时无法判断当前小的数再中间值左边还是右边，例如 [3,1,3,3,3] 和 [3,3,3,1,3]就无法判断1的位置。所以采用左移的方式逐步确定
			low++
		} else {
			high = mid
		}
	}
	return nums[low]
}
