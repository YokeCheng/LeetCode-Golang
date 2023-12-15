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
// getIntersectionNode 返回两个链表的交点节点
// 参数 headA 和 headB 分别是两个链表的头节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// 检查链表是否为空
	if headA == nil || headB == nil {
		return nil
	}

	a, b := headA, headB

	for a != b {
		// 在第一次迭代结束时，将指针重置为另一个链表的头节点
		a = a.Next
		b = b.Next

		// 如果 a 和 b 都为空，说明两个链表没有交点
		if a == nil && b == nil {
			return nil
		}

		// 如果 a 到达链表末尾，将其重置为另一个链表的头节点
		if a == nil {
			a = headB
		}

		// 如果 b 到达链表末尾，将其重置为另一个链表的头节点
		if b == nil {
			b = headA
		}
	}

	// 返回交点节点，如果没有交点则返回 nil
	return a
}
