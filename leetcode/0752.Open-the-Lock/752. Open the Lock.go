package leetcode

import "container/list"

// openLock 使用广度优先搜索算法查找从 "0000" 开始到目标字符串 target 的最短旋转次数。
// 如果无法达到目标或者在 deadends 中找到了死锁状态，则返回 -1。
func openLock(deadends []string, target string) int {
	// 将死锁状态存储在一个集合中，以便快速检查
	deadSet := make(map[string]struct{})
	for _, str := range deadends {
		deadSet[str] = struct{}{}
	}

	// 初始化起始状态和目标状态
	s := "0000"
	if s == target {
		return 0
	}

	// 如果起始状态或目标状态在死锁集合中，直接返回 -1
	if _, exist := deadSet[s]; exist {
		return -1
	}
	if _, exist := deadSet[target]; exist {
		return -1
	}

	// 使用两个队列和映射来执行双向广度优先搜索
	d1 := list.New()           // 用于从起始状态向目标状态搜索
	d2 := list.New()           // 用于从目标状态向起始状态搜索
	m1 := make(map[string]int) // 存储起始状态到达每个状态的旋转次数
	m2 := make(map[string]int) // 存储目标状态到达每个状态的旋转次数

	d1.PushBack(s)
	m1[s] = 0
	d2.PushBack(target)
	m2[target] = 0

	// 迭代执行搜索直到两个队列都为空
	for d1.Len() > 0 && d2.Len() > 0 {
		var t = -1
		if d1.Len() <= d2.Len() {
			t = update(d1, m1, m2, deadSet)
		} else {
			t = update(d2, m2, m1, deadSet)
		}
		if t != -1 {
			return t
		}
	}

	// 如果无法达到目标状态，返回 -1
	return -1
}

// update 函数执行广度优先搜索的一步，从当前状态 d 到达相邻状态，并检查死锁状态。
// 返回最短旋转次数，如果无法到达目标则返回 -1。
func update(d *list.List, cur map[string]int, other map[string]int, deadSet map[string]struct{}) int {
	size := d.Len()
	for i := 0; i < size; i++ {
		fv := d.Remove(d.Front()).(string)
		fvBytes := []byte(fv)
		step := cur[fv]

		// 对每个数字旋转进行迭代
		for j := 0; j < 4; j++ {
			for k := -1; k <= 1; k++ {
				if k == 0 {
					continue
				}
				num := int(fvBytes[j] - '0')
				if k == -1 {
					num -= 1
				} else {
					num += 1
				}
				if num == -1 {
					num = 9
				}
				num = num % 10
				clone := make([]byte, 4)
				copy(clone, fvBytes)
				clone[j] = byte(num) + '0'
				s := string(clone)

				// 检查是否已经访问过状态或者状态在死锁集合中
				if _, exist := cur[s]; exist {
					continue
				}
				if _, exist := deadSet[s]; exist {
					continue
				}

				// 如果相邻状态在另一个方向的映射中存在，返回总旋转次数
				if t, exist := other[s]; exist {
					return step + 1 + t
				}

				// 否则更新映射和队列，继续搜索
				cur[s] = step + 1
				d.PushBack(s)
			}
		}
	}

	// 如果无法到达目标状态，返回 -1
	return -1
}
