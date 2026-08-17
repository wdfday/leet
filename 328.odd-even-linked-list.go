/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	headEven := even 

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = headEven
	return head
}

// @lc code=end

func oddEvenList(head *ListNode) *ListNode {

	dummyOdd := &ListNode{}
	dummyEven := &ListNode{}

	curOdd := dummyOdd
	curEven := dummyEven

	for head != nil {
		if head.Val%2 == 0 {
			curEven.Next = head
			curEven = curEven.Next
		} else {
			curOdd.Next = head
			curOdd = curOdd.Next
		}

		head = head.Next
	}

	curEven.Next = nil 
	curOdd.Next = dummyEven.Next

	return dummyOdd.Next
}
