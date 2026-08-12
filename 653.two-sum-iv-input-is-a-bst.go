/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
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
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]struct{})

	var travel func(node *TreeNode) bool
	travel = func(node *TreeNode) bool {
		if node == nil {
			return false
		}

		if _, ok := m[k-node.Val]; ok {
			return true
		} 
		m[node.Val] = struct{}{}
		return travel(node.Left)|| travel(node.Right)
	}

	return travel(root)    
}
// @lc code=end

