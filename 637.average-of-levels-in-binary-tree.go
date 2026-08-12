/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
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
func averageOfLevels(root *TreeNode) []float64 {

	q := []*TreeNode{root}
	res := []float64{}
	for len(q) > 0 {
		cp := q
		q = nil
		sum := 0.0
		for _, n := range cp {
			sum += float64(n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		res = append(res, sum/float64(len(cp)))
	}
	
    return res
}
// @lc code=end

