package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// TreeNode define
type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 深度优先算法 + 前缀和
func pathSum(root *TreeNode, targetSum int) (ans int) {
	preSum := map[int64]int{0: 1}
	var dfs func(*TreeNode, int64)
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		// 当前的前缀和 - targetSum = 某一个前缀和
		// 判断当前值是否存在，存在则+1
		ans += preSum[curr-int64(targetSum)]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
		return
	}
	dfs(root, 0)
	return
}

// 解法一 带缓存 dfs
func pathSum1(root *TreeNode, targetSum int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	return dfs(root, prefixSum, 0, targetSum)
}

func dfs(root *TreeNode, prefixSum map[int]int, cur, sum int) int {
	if root == nil {
		return 0
	}
	cur += root.Val
	cnt := 0
	if v, ok := prefixSum[cur-sum]; ok {
		cnt = v
	}
	prefixSum[cur]++
	cnt += dfs(root.Left, prefixSum, cur, sum)
	cnt += dfs(root.Right, prefixSum, cur, sum)
	prefixSum[cur]--
	return cnt
}

// 解法二
func pathSumIII(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := findPath437(root, sum)
	res += pathSumIII(root.Left, sum)
	res += pathSumIII(root.Right, sum)
	return res
}

// 寻找包含 root 这个结点，且和为 sum 的路径
func findPath437(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == sum {
		res++
	}
	res += findPath437(root.Left, sum-root.Val)
	res += findPath437(root.Right, sum-root.Val)
	return res
}
