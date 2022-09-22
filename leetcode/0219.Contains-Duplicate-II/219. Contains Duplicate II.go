package leetcode

// 哈希表
func containsNearbyDuplicate1(nums []int, k int) bool {
	pos := map[int]int{}
	// 循环slice，放入哈希map中
	for i, num := range nums {
		// 判断哈希表是否存在，存在则表示值相同，满足nums[i] == nums[j]
		// i代表当前索引，p代表哈希表中索引
		if p, ok := pos[num]; ok && i-p <= k {
			return true
		}
		pos[num] = i
	}
	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	record := make(map[int]bool, len(nums))
	for i, n := range nums {
		// 判断值是否存在
		if _, found := record[n]; found {
			return true
		}
		record[n] = true
		// 由于 abs(i-j) <=k, 则表示大于k的部分都可以去除掉
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}
	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	set := map[int]struct{}{}
	for i, num := range nums {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
