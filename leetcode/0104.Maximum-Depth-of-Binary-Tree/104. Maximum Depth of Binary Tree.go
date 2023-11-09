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

// maxDepth 函数计算给定二叉树的最大深度。
func maxDepth(root *TreeNode) int {
	// 如果根节点为空，返回深度 0
	if root == nil {
		return 0
	}

	// 递归计算左子树和右子树的深度
	leftDept := maxDepth(root.Left)
	rightDept := maxDepth(root.Right)

	// 如果左子树为空，返回右子树深度加1
	if root.Left == nil {
		return rightDept + 1
	}

	// 如果右子树为空，返回左子树深度加1
	if root.Right == nil {
		return leftDept + 1
	}

	// 返回左子树和右子树深度的最大值加1，即整个树的深度
	return max(leftDept, rightDept) + 1
}

// max 函数返回两个整数中的最大值。
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
