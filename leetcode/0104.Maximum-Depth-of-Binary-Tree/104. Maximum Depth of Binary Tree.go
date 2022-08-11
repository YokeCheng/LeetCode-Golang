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

/*func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}*/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	return DFS(root, 0)
}

func DFS(root *TreeNode, level int) int {
	if root == nil {
		return level
	}
	level++
	leftLevel, rightLevel := level, level
	if root.Left != nil {
		leftLevel = DFS(root.Left, level)
	}

	if root.Right != nil {
		rightLevel = DFS(root.Right, level)
	}

	if leftLevel > rightLevel {
		return leftLevel
	}
	return rightLevel
}
