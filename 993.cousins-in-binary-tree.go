/*
 * @lc app=leetcode id=993 lang=golang
 *
 * [993] Cousins in Binary Tree
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
func isCousins(root *TreeNode, x int, y int) bool {
	type item struct {
		node   *TreeNode
		parent *TreeNode
	}
	q := []item{{root, nil}}

	for len(q) > 0 {
		var px, py *TreeNode
		hasX, hasY := false, false
		cp := q
		q = nil
		for _, it := range cp {
			node, parent := it.node, it.parent
			if node.Left != nil {
				q = append(q, item{node.Left, node})
			}
			if node.Right != nil {
				q = append(q, item{node.Right, node})
			}
			if node.Val == x {
				hasX, px = true, parent
			}
			if node.Val == y {
				hasY, py = true, parent
			}
		}
		if hasX != hasY {
			return false
		}
		if hasX && hasY {
			return px != py
		}
	}

	return false
}

// @lc code=end

