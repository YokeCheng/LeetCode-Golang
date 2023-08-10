# [378. Kth Smallest Element in a Sorted Matrix](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/)


## 题目

Implement the RandomizedSet class:

RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
You must implement the functions of the class such that each function works in average O(1) time complexity.



Example 1:

Input
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Output
[null, true, false, true, 2, true, false, 2]

Explanation
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.


Constraints:

-231 <= val <= 231 - 1
At most 2 * 105 calls will be made to insert, remove, and getRandom.
There will be at least one element in the data structure when getRandom is called.


## 题目大意

实现RandomizedSet 类：

RandomizedSet() 初始化 RandomizedSet 对象
bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。



示例：

输入
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
输出
[null, true, false, true, 2, true, false, 2]

解释
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。


提示：

-231 <= val <= 231 - 1
最多调用 insert、remove 和 getRandom 函数 2 * 105 次
在调用 getRandom 方法时，数据结构中 至少存在一个 元素。


## 解题思路
```go
type RandomizedSet struct {
    nums    []int       // 存储实际元素值的数组
    indices map[int]int // 以元素值作为键、索引作为值的哈希表
}

// 构造函数，初始化 RandomizedSet 结构体
func Constructor() RandomizedSet {
    return RandomizedSet{[]int{}, map[int]int{}}
}

// 插入元素到集合中
func (rs *RandomizedSet) Insert(val int) bool {
    if _, ok := rs.indices[val]; ok { // 判断元素是否已经存在于哈希表中
        return false // 元素已存在，插入失败
    }
    rs.indices[val] = len(rs.nums) // 在哈希表中添加元素值和对应的索引
    rs.nums = append(rs.nums, val) // 在数组末尾添加元素值
    return true // 插入成功
}

// 从集合中删除元素
func (rs *RandomizedSet) Remove(val int) bool {
    id, ok := rs.indices[val] // 获取要删除元素在哈希表中的索引
    if !ok { // 判断元素是否存在于哈希表中
        return false // 元素不存在，删除失败
    }
    last := len(rs.nums) - 1 // 获取数组最后一个元素的索引
    rs.nums[id] = rs.nums[last] // 用数组最后一个元素替换要删除的元素
    rs.indices[rs.nums[id]] = id // 更新最后一个元素在哈希表中的索引
    rs.nums = rs.nums[:last] // 删除数组最后一个元素
    delete(rs.indices, val) // 删除哈希表中的旧键值对
    return true // 删除成功
}

// 获取集合中的随机元素
func (rs *RandomizedSet) GetRandom() int {
    return rs.nums[rand.Intn(len(rs.nums))] // 随机获取数组中的元素值
}
```


