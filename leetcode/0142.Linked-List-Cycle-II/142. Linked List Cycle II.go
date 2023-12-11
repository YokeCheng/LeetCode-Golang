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
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// 如果链表为空或者链表的头部为空，直接返回nil
	if head == nil || head.Next == nil {
		return nil
	}

	// 快指针从头节点开始，慢指针从头节点开始
	fast := head
	slow := head

	// 当慢指针和快指针都不为空，且快指针的下一个节点不为空时，进行循环
	for slow != nil && fast != nil && fast.Next != nil {
		// 快指针移动两步
		fast = fast.Next.Next
		// 慢指针移动一步
		slow = slow.Next
		// 如果快指针追上了慢指针，说明链表中存在环
		if fast == slow {
			// 重置快指针到链表头部
			fast = head
			// 当快指针和慢指针相遇时，重新开始循环，直到快指针和慢指针相遇为止
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			// 当快指针和慢指针相遇时，返回相遇节点，这个节点是环的入口。
			return fast
		}
	}

	// 如果快指针追上慢指针的情况没有发生，说明链表中不存在环，返回nil。
	return nil
}
