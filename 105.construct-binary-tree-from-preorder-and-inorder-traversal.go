/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	pos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		pos[v] = i
	}

	preIdx := 0
	var dfs func(l, r int) *TreeNode
	dfs = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		rootVal := preorder[preIdx]
		preIdx++

		mid := pos[rootVal]
		root := &TreeNode{Val: rootVal}

		root.Left = dfs(l, mid-1)
		root.Right = dfs(mid+1, r)

		return root
	}

	return dfs(0, len(inorder)-1)
}

// @lc code=end

