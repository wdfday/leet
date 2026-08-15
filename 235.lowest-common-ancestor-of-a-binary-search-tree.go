/*
 * @lc app=leetcode id=235 lang=golang
 *
 * [235] Lowest Common Ancestor of a Binary Search Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	
	if q.Val < root.Val && p.Val < root.Val   {
		return lowestCommonAncestor(root.Left, p, q)
	} else if q.Val > root.Val && p.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} 
	return root
}
// @lc code=end

