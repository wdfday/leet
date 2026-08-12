/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}

	first := head

	for first != nil && first.Next != nil {
		if first.Val == first.Next.Val {
			first.Next = first.Next.Next
			continue
		}
		first = first.Next
	}
	return dummy.Next
}

// @lc code=end

// @lc code=end

