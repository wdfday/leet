/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	left := sortedArrayToBST(nums[:mid])
	right := sortedArrayToBST(nums[mid+1:])

	res := &TreeNode{Val: nums[mid], Left: left, Right: right}

	return res

}

// @lc code=end

