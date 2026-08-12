/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {

	dummy := &ListNode{0, head}

	first := dummy
	for first.Next != nil {
		if first.Next.Val == val {
			first.Next = first.Next.Next
			continue
		}
		first = first.Next
	}
	return dummy.Next

}

// @lc code=end

