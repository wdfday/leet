/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
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
func findMode(root *TreeNode) []int {
	var res []int
	var prev *TreeNode
	count := 0
	maxCount := 0

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		if prev != nil && prev.Val == node.Val {
			count++
		} else {
			count = 1
		}

		if count > maxCount {
			maxCount = count
			res = []int{node.Val}
		} else if count == maxCount {
			res = append(res, node.Val)
		}

		prev = node
		inorder(node.Right)
	}

	inorder(root)
	return res
}

// @lc code=end

