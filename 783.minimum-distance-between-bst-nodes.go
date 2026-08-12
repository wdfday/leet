/*
 * @lc app=leetcode id=783 lang=golang
 *
 * [783] Minimum Distance Between BST Nodes
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
func minDiffInBST(root *TreeNode) int {
	q := []int{}
	res := math.MaxInt
	var preOrder func(root *TreeNode) 
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		preOrder(root.Left)
		q = append(q, root.Val)
		if len(q) >= 2 {
			d := q[len(q) - 1] - q[len(q) - 2]
			if d < res {
				res = d
			}
		}
		preOrder(root.Right)
	}
	preOrder(root)
    return res
}
// @lc code=end

