/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
func binaryTreePaths(root *TreeNode) []string {

	if root == nil {
		return nil
	}
	res := []string{}
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		if path == "" {
			path = strconv.Itoa(node.Val)
		} else {
			path += "->" + strconv.Itoa(node.Val)
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, "")
	return res

}

// @lc code=end

