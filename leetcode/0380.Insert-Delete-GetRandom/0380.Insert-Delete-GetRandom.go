package leetcode

import "math/rand"

type RandomizedSet struct {
	nums    []int       // 存储实际元素值的数组
	indices map[int]int // 以元素值作为键、索引作为值的哈希表
}

// Constructor 构造函数，初始化 RandomizedSet 结构体
func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

// Insert 插入元素到集合中
func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.indices[val]; ok { // 判断元素是否已经存在于哈希表中
		return false // 元素已存在，插入失败
	}
	rs.indices[val] = len(rs.nums) // 在哈希表中添加元素值和对应的索引
	rs.nums = append(rs.nums, val) // 在数组末尾添加元素值
	return true                    // 插入成功
}

// Remove 从集合中删除元素
func (rs *RandomizedSet) Remove(val int) bool {
	id, ok := rs.indices[val] // 获取要删除元素在哈希表中的索引
	if !ok {                  // 判断元素是否存在于哈希表中
		return false // 元素不存在，删除失败
	}
	last := len(rs.nums) - 1     // 获取数组最后一个元素的索引
	rs.nums[id] = rs.nums[last]  // 用数组最后一个元素替换要删除的元素
	rs.indices[rs.nums[id]] = id // 更新最后一个元素在哈希表中的索引
	rs.nums = rs.nums[:last]     // 删除数组最后一个元素
	delete(rs.indices, val)      // 删除哈希表中的旧键值对
	return true                  // 删除成功
}

// GetRandom 获取集合中的随机元素
func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))] // 随机获取数组中的元素值
}
