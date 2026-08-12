/*
 * @lc app=leetcode id=671 lang=golang
 *
 * [671] Second Minimum Node In a Binary Tree
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
func findSecondMinimumValue(root *TreeNode) int {
	res := math.MaxInt
	var dfs func(node *TreeNode) 
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Val < res && node.Val != root.Val {
			res = node.Val
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
dfs(root)
	if res == math.MaxInt {
		return -1
	}
    return res
}
// @lc code=end

func findSecondMinimumValue(root *TreeNode) int {
	s := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		s = append(s, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	slices.Sort(s)
	for _, v := range s {
		if v > s[0] {
			return v
		}
	}

    return -1
}