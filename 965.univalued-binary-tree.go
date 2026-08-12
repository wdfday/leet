/*
 * @lc app=leetcode id=965 lang=golang
 *
 * [965] Univalued Binary Tree
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
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if root.Val != root.Left.Val {
			return false
		}
	}
	if root.Right != nil {
		if root.Val != root.Right.Val {
			return false
		}
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

// @lc code=end

