/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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
func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	res := []int{}
	res = append(left, root.Val)
	res = append(res, right...)
	return res
}

// @lc code=end

