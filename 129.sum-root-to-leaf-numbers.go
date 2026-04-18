/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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
func sumNumbers(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var dfs func(node *TreeNode, path int) int
	dfs = func(node *TreeNode, path int) int {
		if node == nil {
			return 0
		}

		path = path*10 + node.Val

		if node.Left == nil && node.Right == nil {
			return path
		}

		left := dfs(node.Left, path)
		right := dfs(node.Right, path)

		return left + right
	}

	return dfs(root, 0)
}

// @lc code=end

