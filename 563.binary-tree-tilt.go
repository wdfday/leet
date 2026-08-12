/*
 * @lc app=leetcode id=563 lang=golang
 *
 * [563] Binary Tree Tilt
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
func findTilt(root *TreeNode) int {
	res := 0

	var postOrder func(node *TreeNode) int
	postOrder = func(node *TreeNode) int{
		if node == nil {
			return 0
		}
		l := postOrder(node.Left)
		r := postOrder(node.Right)

		res += abs(l - r)
		sum := node.Val + l + r
		return sum
	}
	postOrder(root)
    return res
}

func abs(a int) int {
	if a < 0 {
		a = - a
	}
	return a
}
// @lc code=end

