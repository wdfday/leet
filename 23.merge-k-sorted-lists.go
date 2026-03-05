/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		var mergedLists []*ListNode
		for i := 0; i < len(lists); i += 2 {
			if i+1 < len(lists) {
				mergedLists = append(mergedLists, mergeTwoLists(lists[i], lists[i+1]))
			} else {
				mergedLists = append(mergedLists, lists[i])
			}
		}
		lists = mergedLists
	}
	return lists[0]
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			current.Next = list2
			list2 = list2.Next
		} else {
			current.Next = list1
			list1 = list1.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	return dummy.Next
}

// @lc code=end

// /*
//  * @lc app=leetcode id=23 lang=golang
//  *
//  * [23] Merge k Sorted Lists
//  */

// import "container/heap"

// // @lc code=start
// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */

// type NodeHeap []*ListNode

// func (h NodeHeap) Len() int            { return len(h) }
// func (h NodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
// func (h NodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
// func (h *NodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
// func (h *NodeHeap) Pop() interface{} {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     *h = old[:n-1]
//     return x
// }

// func mergeKLists(lists []*ListNode) *ListNode {
//     h := &NodeHeap{}
//     heap.Init(h)

//     for _, node := range lists {
//         if node != nil {
//             heap.Push(h, node)
//         }
//     }

//     var dummy ListNode
//     cur := &dummy

//     for h.Len() > 0 {
//         minNode := heap.Pop(h).(*ListNode)
//         cur.Next = minNode
//         cur = cur.Next

//         if minNode.Next != nil {
//             heap.Push(h, minNode.Next)
//         }
//     }

//     return dummy.Next
// }

