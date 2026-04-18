/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}

	var traversal func(node *TreeNode, target int, cur []int)
	traversal = func(node *TreeNode, target int, cur []int) {
		if node == nil {
			return
		}

		cur = append(cur, node.Val)

		if node.Left == nil && node.Right == nil && target == node.Val {
			res = append(res, append([]int{}, cur...))
			return
		} else {
			traversal(node.Left, target-node.Val, cur)
			traversal(node.Right, target-node.Val, cur)
		}
	}

	traversal(root, targetSum, []int{})
	return res
}

// @lc code=end

