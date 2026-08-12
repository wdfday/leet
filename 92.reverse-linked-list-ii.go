/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left >= right {
		return head
	}

	dummy := &ListNode{Next: head}
	pre := dummy
	// move `pre` to the node before position `left`
	for i := 0; i < left-1; i++ {
		if pre == nil {
			return head
		}
		pre = pre.Next
	}

	// `curr` is the first node to be reversed
	curr := pre.Next
	// perform head-insertion reversal for (right-left) times
	for i := 0; i < right-left; i++ {
		if curr == nil || curr.Next == nil {
			break
		}
		nxt := curr.Next
		// remove nxt
		curr.Next = nxt.Next
		// insert nxt after pre
		nxt.Next = pre.Next
		pre.Next = nxt
	}

	return dummy.Next
}

// @lc code=end

