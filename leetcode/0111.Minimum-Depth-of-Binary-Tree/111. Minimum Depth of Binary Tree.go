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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	// 当左子树为空时，返回右子树的深度加1
	if root.Left == nil {
		return rightDepth + 1
	}

	// 当右子树为空时，返回左子树的深度加1
	if root.Right == nil {
		return leftDepth + 1
	}

	// 否则，返回较小深度的子树深度加1
	return min(leftDepth, rightDepth) + 1
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
