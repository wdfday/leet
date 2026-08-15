/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := []*Node{root}

	for len(q) > 0 {
		cur := q
		q = []*Node{}

		for i := 0; i < len(cur); i++ {
			if i < len(cur) - 1 {
				cur[i].Next = cur[i+1]
			}

			if cur[i].Left != nil {
				q = append(q, cur[i].Left)
			}
			if cur[i].Right != nil {
				q = append(q, cur[i].Right)
			}
		}
	}

	return root
}
// @lc code=end

