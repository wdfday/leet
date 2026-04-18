/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	leftToRight := true

	for len(q) > 0 {
		size := len(q)
		level := make([]int, size)

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			idx := i
			if !leftToRight {
				idx = size - 1 - i
			}
			level[idx] = node.Val

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, level)
		leftToRight = !leftToRight
	}

	return res
}

// @lc code=end

