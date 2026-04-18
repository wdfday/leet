/*
 * @lc app=leetcode id=111 lang=golang
 *
 * [111] Minimum Depth of Binary Tree
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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type NodeDepth struct {
		node  *TreeNode
		depth int
	}

	q := []NodeDepth{{root, 1}}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		n := cur.node
		d := cur.depth

		if n.Left == nil && n.Right == nil {
			return d
		}

		if n.Left != nil {
			q = append(q, NodeDepth{n.Left, d + 1})
		}
		if n.Right != nil {
			q = append(q, NodeDepth{n.Right, d + 1})
		}
	}

	return 0
}

// func minDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		return 1
// 	}

// 	left := minDepth(root.Left)
// 	right := minDepth(root.Right)

// 	if left != 0 && right != 0 {
// 		return min(left, right) + 1
// 	}

// 	if left == 0 {
// 		return right + 1
// 	}

// 	return left + 1

// }

// @lc code=end

