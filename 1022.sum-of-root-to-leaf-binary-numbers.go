/*
 * @lc app=leetcode id=1022 lang=golang
 *
 * [1022] Sum of Root To Leaf Binary Numbers
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
func sumRootToLeaf(root *TreeNode) int {
	res := 0
	var preOrder func(node *TreeNode, pref int) 
	preOrder = func(node *TreeNode, pref int) {
		if node == nil {
			return
		}
		pref = pref * 2 + node.Val

		if node.Left == nil && node.Right == nil {
			res += pref
		} else if node.Left == nil {
			preOrder(node.Right, pref)
		} else if node.Right == nil {
			preOrder(node.Left, pref)
		} else {
			preOrder(node.Right, pref)
			preOrder(node.Left, pref)

		}
	}

	preOrder(root, 0)
	return res
    
}
// @lc code=end

