/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Tìm độ dài và tail
	n := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}

	k %= n
	if k == 0 {
		return head
	}

	// Nối thành vòng tròn
	tail.Next = head

	// newTail ở vị trí n-k-1 (0-index từ head)
	steps := n - k - 1
	newTail := head
	for i := 0; i < steps; i++ {
		newTail = newTail.Next
	}

	// Cắt vòng
	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}

// @lc code=end

