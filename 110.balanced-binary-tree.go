/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	var height func(node *TreeNode) int

	// Trả về chiều cao; trả -1 nếu subtree mất cân bằng.
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := height(node.Left)
		if left == -1 {
			return -1
		}

		right := height(node.Right)
		if right == -1 {
			return -1
		}

		if left-right > 1 || right-left > 1 {
			return -1
		}

		if left > right {
			return left + 1
		}
		return right + 1
	}

	return height(root) != -1
}

// @lc code=end

