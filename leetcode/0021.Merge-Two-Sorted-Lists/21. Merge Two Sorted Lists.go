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

func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil { // 如果list1为空，则直接返回list2
		return list2
	}
	if list2 == nil { // 如果list2为空，则直接返回list1
		return list1
	}

	// 在这里我们可以直接在list1和list2上操作，不需要引入新的变量。
	if list1.Val < list2.Val { // 如果list1当前节点较小，则将其插入到list1后面
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else { // 否则将list2当前节点插入到list1后面
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy ListNode // 创建一个虚拟头结点，使链表可以反向插入

	curr := &dummy // curr 指向虚拟头结点

	for list1 != nil && list2 != nil { // 当两个链表都不为空时，进行比较并插入较小的节点
		if list1.Val < list2.Val { // 如果list1当前节点较小，则将其插入到curr之后
			curr.Next = list1  // 将curr后移一位
			list1 = list1.Next // 将list1后移一位
		} else { // 否则将list2当前节点插入到curr之后
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next // 将curr后移一位
	}

	// 插入最后一个不为空的链表的所有节点（如果有）
	for ; list1 != nil; list1 = list1.Next { // 如果list1还有剩余节点，则将所有剩余节点插入到curr之后
		curr.Next = list1 // 将curr后移一位
		curr = curr.Next  // 将curr后移一位
	}
	for ; list2 != nil; list2 = list2.Next { // 如果list2还有剩余节点，则将所有剩余节点插入到curr之后
		curr.Next = list2 // 将curr后移一位
		curr = curr.Next  // 将curr后移一位
	}

	return dummy.Next // 返回合并后的链表的头节点
}
