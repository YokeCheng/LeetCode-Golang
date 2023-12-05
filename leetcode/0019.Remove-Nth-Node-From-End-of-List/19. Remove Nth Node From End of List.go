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

// 解法一
// removeNthFromEnd 从链表末尾删除倒数第 N 个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 使用虚拟头结点简化边界情况处理
	dummyHead := &ListNode{Next: head}
	// preSlow 用于记录要删除节点的前一个节点
	// slow 和 fast 用于遍历链表
	preSlow, slow, fast := dummyHead, head, head

	// 使用快慢指针，fast 先走 N 步，然后 slow 和 fast 一起走
	for fast != nil {
		// 当 n 小于等于 0 时，preSlow 和 slow 开始同步移动
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		// 每遍历一个节点，n 减一
		n--
		// fast 继续向前移动
		fast = fast.Next
	}

	// 删除倒数第 N 个节点
	preSlow.Next = slow.Next

	// 返回虚拟头结点的下一个节点，即处理后的链表头
	return dummyHead.Next
}

// 解法二
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	current := head
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	if n > len {
		return head
	}
	if n == len {
		current := head
		head = head.Next
		current.Next = nil
		return head
	}
	current = head
	i := 0
	for current != nil {
		if i == len-n-1 {
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		current = current.Next
	}
	return head
}
