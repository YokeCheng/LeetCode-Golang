package leetcode

import "github.com/halfrost/LeetCode-Go/structures"

//Definition for singly-linked list.
func reversePrint(head *structures.ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint(head.Next), head.Val)
}
