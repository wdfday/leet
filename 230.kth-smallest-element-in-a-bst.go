/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
func kthSmallest(root *TreeNode, k int) int {
	res := -1

	var preOrder func(node *TreeNode, p int) int 
	preOrder = func(node *TreeNode, p int) int {
		if node == nil {
			return 0
		}

		lr := preOrder(node.Left, p)
		if p + lr == k - 1 {
			res = node.Val
		}
		r := preOrder(node.Right, p + lr + 1)
		return lr + r + 1
	}

	preOrder(root, 0)
	return res
}
// @lc code=end

