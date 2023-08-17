package leetcode

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	// 从左到右遍历，确保右边评分高的孩子得到更多糖果
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] += candies[i-1] + 1
		}
	}
	// 从右到左遍历，确保左边评分高的孩子得到更多糖果
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	total := 0
	for _, candy := range candies {
		total += candy + 1
	}
	return total
}

func candy2(ratings []int) int {
	n := len(ratings)
	ans, inc, dec, pre := 1, 1, 0, 1

	// 从第二个孩子开始遍历
	for i := 1; i < n; i++ {
		// 如果当前孩子的评分大于等于前一个孩子的评分
		if ratings[i] >= ratings[i-1] {
			dec = 0 // 递减计数器重置为0
			if ratings[i] == ratings[i-1] {
				pre = 1 // 如果评分相等，连续评分为1
			} else {
				pre++ // 否则连续评分递增
			}
			ans += pre // 分配糖果数量为连续评分之和
			inc = pre  // 递增计数器更新
		} else {
			dec++ // 孩子评分递减，递减计数器增加
			if dec == inc {
				dec++ // 如果递减计数器与递增计数器相等，递减计数器增加
			}
			ans += dec // 分配糖果数量为递减计数器的值
			pre = 1    // 连续评分重置为1
		}
	}
	return ans
}
