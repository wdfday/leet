/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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

func generateTrees(n int) []*TreeNode {
	var gen func(l, r int) []*TreeNode
	gen = func(l, r int) []*TreeNode {
		res := []*TreeNode{}

		if l > r {
			return append(res, nil) // base: slot cho nil
		}

		for i := l; i <= r; i++ {
			left := gen(l, i-1)
			right := gen(i+1, r)

			for _, lNode := range left {
				for _, rNode := range right {
					root := &TreeNode{Val: i}
					root.Left = lNode
					root.Right = rNode
					res = append(res, root)
				}
			}
		}

		return res
	}

	return gen(1, n)
}

// @lc code=end

