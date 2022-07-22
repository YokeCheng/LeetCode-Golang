package leetcode

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	// j 判断第一个0的位置
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// i = j表示不需要调换位置
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
}

func moveZeroes1(nums []int) {
	numsLen := len(nums)
	if numsLen == 0 {
		return
	}

	j := 0
	for i := 0; i < numsLen; i++ {
		if nums[i] != 0 {
			if i != j {
				nums[i], nums[j] = nums[i], nums[j]
			}
			j++
		}
	}
}
