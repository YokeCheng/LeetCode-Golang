package leetcode

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	var mapA = map[int]int{}
	for i, num := range nums {
		if index, ok := mapA[num]; ok {
			return []int{index, i}
		}
		mapA[target-num] = i
	}
	return nil
}
