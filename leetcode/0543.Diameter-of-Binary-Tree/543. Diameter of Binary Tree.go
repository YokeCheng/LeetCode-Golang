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

// max 函数接受两个整数参数，返回它们之间的较大值。
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// diameterOfBinaryTree 函数接受一个指向二叉树根节点的指针，并返回二叉树的直径。
func diameterOfBinaryTree(root *TreeNode) int {
	// 初始化最大直径为0
	maxDiameter := 0
	// 调用 getMaxDiameter 函数来计算最大直径
	getMaxDiameter(root, &maxDiameter)
	// 返回最大直径
	return maxDiameter
}

// getMaxDiameter 函数接受一个二叉树节点和一个整数指针，用于计算以该节点为根的子树的最大直径。
// 它返回以该节点为根的子树的深度，并通过整数指针更新最大直径。
func getMaxDiameter(root *TreeNode, maxDiameter *int) int {
	// 如果节点为空，返回深度0
	if root == nil {
		return 0
	}
	// 递归计算左子树和右子树的深度
	left := getMaxDiameter(root.Left, maxDiameter)
	right := getMaxDiameter(root.Right, maxDiameter)
	// 更新最大直径，考虑经过当前节点的直径和已知最大直径的较大值
	*maxDiameter = max(left+right, *maxDiameter)
	// 返回当前子树的深度，加1表示包括当前节点
	return max(left, right) + 1
}
