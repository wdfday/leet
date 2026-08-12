/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {

	var min int64 = -1 << 63
	var max int64 = 1<<63 - 1

	return validBranch(root, min, max)

}

func validBranch(node *TreeNode, left, right int64) bool {
	if node == nil {
		return true
	}

	v := int64(node.Val)
	if v <= left || v >= right {
		return false
	}

	return validBranch(node.Left, left, v) && validBranch(node.Right, v, right)

}

// @lc code=end

