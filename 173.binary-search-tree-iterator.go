/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
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
type BSTIterator struct {
    s []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
	x := []*TreeNode{}
	for root != nil {
		x = append(x, root)
		root = root.Left
	}
	return BSTIterator{x}
}


func (this *BSTIterator) Next() int {
	next := 0
	m := len(this.s)
	if m == 0 {
		return next
	}
    last := this.s[m - 1]

	next = last.Val
	this.s = this.s[:m - 1]
	if last.Right != nil {
		last = last.Right 
		for last != nil {
			this.s = append(this.s, last)
			last = last.Left
		}
	}
	return next
}



func (this *BSTIterator) HasNext() bool {
    return len(this.s) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

