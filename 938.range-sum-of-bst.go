/*
 * @lc app=leetcode id=938 lang=golang
 *
 * [938] Range Sum of BST
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
func rangeSumBST(root *TreeNode, low int, high int) int {
	res := 0
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return 
		}

		if node.Val < low {
			inorder(node.Right)
		} else if node.Val > high {
			inorder(node.Left)
		} else {
			res += node.Val
			inorder(node.Left)
			inorder(node.Right)
		}
	}
	inorder(root)
    return res
}
// @lc code=end

