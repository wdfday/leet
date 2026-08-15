/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
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
	dummy := &ListNode{Next:head,}
	cur := dummy
	nxt := head
	for nxt != nil {
		if nxt.Next != nil && nxt.Val == nxt.Next.Val {
			newNxt := nxt.Next
			for newNxt != nil && nxt.Val == newNxt.Val {
				newNxt = newNxt.Next
			}
			cur.Next = newNxt
			nxt = newNxt
		} else {
			cur = nxt
			nxt = nxt.Next
		}
	}
    return dummy.Next
}
// @lc code=end

