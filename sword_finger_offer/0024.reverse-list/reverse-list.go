package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

//Definition for singly-linked list.
func reverseList1(head *structures.ListNode) *structures.ListNode {
	var prev *structures.ListNode
	for head != nil {
		prev = &structures.ListNode{
			Val:  head.Val,
			Next: prev,
		}
		head = head.Next
	}
	return prev
}
