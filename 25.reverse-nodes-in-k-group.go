/*
 * @lc app=leetcode id=25 lang=golang
 *
 * [25] Reverse Nodes in k-Group
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	dummy := &ListNode{Next: head}
	prevGroupEnd := dummy

	for {
		kthNode := prevGroupEnd
		for i := 0; i < k && kthNode != nil; i++ {
			kthNode = kthNode.Next
		}
		if kthNode == nil {
			break
		}

		groupStart := prevGroupEnd.Next
		nextGroupStart := kthNode.Next

		// Reverse the current group
		prev, curr := nextGroupStart, groupStart
		for curr != nextGroupStart {
			temp := curr.Next
			curr.Next = prev
			prev = curr
			curr = temp
		}

		prevGroupEnd.Next = kthNode
		prevGroupEnd = groupStart
	}

	return dummy.Next
}

// @lc code=end

