package leetcode

// 解法一 时间复杂度 O(n)，空间复杂度 O(1)
func rotate1(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

// 解法二 时间复杂度 O(n)，空间复杂度 O(n)
// 新设置一个空间，将值存入
func rotate2(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

// 解法3 时间复杂度 O(n)，空间复杂度 O(1)
// 设置一个临时空间，将要替换的值存入
func rotate(nums []int, k int) {
	temp, next, start := nums[0], 0, 0
	for i := 0; i < len(nums); i++ {
		next = (next + k) % len(nums)
		temp, nums[next] = nums[next], temp
		// 当回到起始index，则表示当前替换已成环，需要进行下一步替换
		if next == start {
			start++
			next = start
			temp = nums[start]
		}
	}
}

func rotate5(nums []int, k int) {
	k1 := k % len(nums)

	newNums := append(nums[len(nums)-k1:], nums[:len(nums)-k1]...)
	copy(nums, newNums)
}

// 解法4 时间复杂度 O(n)，空间复杂度 O(1)
// 和解法3相同，只是使用了最大公约数判断，更好
func rotate4(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

// 取最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
