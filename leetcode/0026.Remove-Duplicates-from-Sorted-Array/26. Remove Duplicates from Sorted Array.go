package leetcode

// 解法一
// 通过双指针，last指的是不重复的当前值，内循环finder找到下一个和last不同的值，
// 然后将finder值和last的下一个值替换
func removeElement3(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] == nums[j] {
			continue
		}
		nums[i+1] = nums[j]
		i++
	}
	return i + 1
}

// 解法二
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	length := len(nums)
	lastNum := nums[length-1]
	i := 0
	for i = 0; i < length-1; i++ {
		if nums[i] == lastNum {
			break
		}
		if nums[i+1] == nums[i] {
			removeElement1(nums, i+1, nums[i])
			// fmt.Printf("此时 num = %v length = %v\n", nums, length)
		}
	}
	return i + 1
}

func removeElement1(nums []int, start, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := start
	for i := start; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
	return j
}
