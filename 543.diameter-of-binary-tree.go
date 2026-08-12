/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	var diameter func(node *TreeNode) int 
	diameter = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		l := diameter(node.Left)
		r := diameter(node.Right)
		if l + r > res {
			res = r + l
		}


		return 1 + max(l, r)

	}
    diameter(root)
	return res
}
// @lc code=end

