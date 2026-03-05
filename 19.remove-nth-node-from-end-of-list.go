/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{0, head}
	first := dummy
	second := dummy

	// Move first pointer n+1 steps ahead
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// Move both pointers until first reaches the end
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// Remove the nth node from the end
	second.Next = second.Next.Next

	return dummy.Next
}

// @lc code=end

