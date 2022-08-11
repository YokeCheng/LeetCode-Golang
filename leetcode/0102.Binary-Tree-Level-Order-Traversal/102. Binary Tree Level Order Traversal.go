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

// 解法一 BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[l:]
		res = append(res, tmp)
	}
	return res
}

// 解法二 DFS
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	var dfsLevel func(node *TreeNode, level int)
	dfsLevel = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{node.Val})
		} else {
			res[level] = append(res[level], node.Val)
		}
		dfsLevel(node.Left, level+1)
		dfsLevel(node.Right, level+1)
	}
	dfsLevel(root, 0)
	return res
}

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	res := make([][]int, 0)
	level := 1
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if level%2 == 1 {
				tmp = append(tmp, queue[i].Val)
			} else {
				tmp = append([]int{queue[i].Val}, tmp...)
			}
		}
		queue = queue[l:]
		res = append(res, tmp)
		level++
	}
	return res
}

func isStraight(nums []int) bool {
	var mapN [14]bool
	min, max := 14, 0

	for _, num := range nums {
		if num == 0 {
			continue
		}
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		if mapN[num] {
			return false
		}
		mapN[num] = true
	}

	if max-min <= 4 {
		return true
	}

	return false
}
