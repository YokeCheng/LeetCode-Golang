package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// middleNode 返回链表的中间节点。如果链表有偶数个节点，则返回中间两个节点的第一个节点。
func middleNode(head *ListNode) *ListNode {
	// 使用两个指针，一个慢指针（slow）和一个快指针（fast），初始位置都在链表头部
	slow, fast := head, head

	// 遍历链表，快指针每次移动两步，慢指针每次移动一步
	for fast != nil && fast.Next != nil {
		// 快指针先移动一步
		fast = fast.Next
		// 如果快指针还能再移动一步，再移动一步
		if fast != nil {
			fast = fast.Next
		}
		// 慢指针每次移动一步
		slow = slow.Next
	}

	// 当快指针到达链表末尾时，慢指针所在位置即为中间节点的位置
	return slow
}
