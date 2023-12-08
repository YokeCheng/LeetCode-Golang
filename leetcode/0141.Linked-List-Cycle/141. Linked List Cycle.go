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

// hasCycle 检查链表是否有循环
func hasCycle(head *ListNode) bool {
	// 初始化两个指针，fast（快指针）和slow（慢指针），它们都从链表的头部开始
	fast := head
	slow := head

	// 遍历链表，直到fast指针到达链表末尾或检测到循环
	for fast != nil && fast.Next != nil {
		// fast指针每次移动两步
		fast = fast.Next.Next
		// slow指针每次移动一步
		slow = slow.Next

		// 检查fast和slow指针是否相遇，表示存在循环
		if fast == slow {
			// 检测到循环，返回true
			return true
		}
	}

	// 如果循环完成而没有检测到循环，则返回false
	return false
}
