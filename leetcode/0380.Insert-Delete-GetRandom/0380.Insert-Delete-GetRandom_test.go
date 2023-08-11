package leetcode

import (
	"testing"
)

func Test_Problem378(t *testing.T) {
	randomizedSet := Constructor()
	randomizedSet.Insert(1)   // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
	randomizedSet.Remove(2)   // 返回 false ，表示集合中不存在 2 。
	randomizedSet.Insert(2)   // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
	randomizedSet.GetRandom() // getRandom 应随机返回 1 或 2 。
	randomizedSet.Remove(1)   // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
	randomizedSet.Insert(2)   // 2 已在集合中，所以返回 false 。
	randomizedSet.GetRandom() // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
}
