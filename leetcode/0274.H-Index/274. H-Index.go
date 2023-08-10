package leetcode

import "sort"

func hIndex(citations []int) int {
	//倒序排序
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	n := len(citations)
	// 当前值小于下标+1(实际论文数)，则表示不符合要求，那符合要求为上一个论文数
	for i := 0; i < n; i++ {
		if citations[i] < i+1 {
			return i
		}
	}
	return n
}

// 解法一
func hIndex2(citations []int) int {

	// n为论文总数
	n := len(citations)

	// count数组统计每个引用次数的论文数量
	// 长度为n+1,以包含引用次数大于n的情况
	count := make([]int, n+1)
	for _, c := range citations {
		// 如果引用次数大于等于n
		// 则计数到count数组最后一个元素中
		if c >= n {
			count[n]++
			// 否则统计到对应引用次数的计数桶中
		} else {
			count[c]++
		}
	}

	// k统计累积的论文数量
	k := 0

	// 从后向前遍历count数组
	for i := n; i >= 0; i-- {

		// 累加当前引用次数的论文数量
		k += count[i]

		// 如果累积论文数量达到当前引用次数
		// 即满足h指数的定义
		if k >= i {

			// 返回当前引用次数i作为h指数
			return i
		}
	}

	// 如果没有找到满足的i
	// 返回0
	return 0
}

// 解法二
func hIndex1(citations []int) int {
	quickSort164(citations, 0, len(citations)-1)
	hIndex := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= len(citations)-i {
			hIndex++
		} else {
			break
		}
	}
	return hIndex
}

func quickSort164(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition164(a, lo, hi)
	quickSort164(a, lo, p-1)
	quickSort164(a, p+1, hi)
}

func partition164(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
