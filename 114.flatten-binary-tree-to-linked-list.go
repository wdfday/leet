/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
func flatten(root *TreeNode) {
	cur := root
	for cur != nil {
		if cur.Left != nil {
			// Tìm node ngoài cùng bên phải của nhánh trái
			pre := cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}

			// Gắn nhánh phải cũ vào cuối nhánh trái
			pre.Right = cur.Right

			// Đưa nhánh trái sang phải
			cur.Right = cur.Left
			cur.Left = nil
		}
		// Tiến sang node tiếp theo trên "linked list"
		cur = cur.Right
	}
}

// @lc code=end

