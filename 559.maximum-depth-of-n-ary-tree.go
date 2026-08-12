/*
 * @lc app=leetcode id=559 lang=golang
 *
 * [559] Maximum Depth of N-ary Tree
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	res := 0
	var dfs func(node *Node, d int) 
	dfs = func(node *Node, d int) {
		if node == nil {
			return
		}

		d++ 
		if d > res {
			res = d
		}

		for _, child := range node.Children {
			dfs(child, d)
		}
		return
	}
	dfs(root, 0)
    return res
}
// @lc code=end

