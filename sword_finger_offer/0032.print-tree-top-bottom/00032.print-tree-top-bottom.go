package leetcode

import (
	"container/list"
	"github.com/halfrost/LeetCode-Go/structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode define
type TreeNode = structures.TreeNode

// 广度优先排序
func levelOrder(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		// 取出节点
		tmp := queue.Remove(queue.Front()).(*TreeNode)

		//放入ans中
		ans = append(ans, tmp.Val)

		// 将下层左右节点塞入进去
		if tmp.Left != nil {
			queue.PushBack(tmp.Left)
		}

		if tmp.Right != nil {
			queue.PushBack(tmp.Right)
		}
	}
	return ans
}

// 32-2题
func levelOrder32_2(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() != 0 {
		// 判断数量，保证只处理到一层的数据
		curLen := queue.Len()
		curAns := []int{}
		for i := 0; i < curLen; i++ {
			//取出数据
			tmp := queue.Remove(queue.Front()).(*TreeNode)

			curAns = append(curAns, tmp.Val)
			// 将下层左右节点塞入进去
			if tmp.Left != nil {
				queue.PushBack(tmp.Left)
			}

			if tmp.Right != nil {
				queue.PushBack(tmp.Right)
			}
		}
		ans = append(ans, curAns)
	}
	return ans
}
