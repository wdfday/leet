/*
 * @lc app=leetcode id=589 lang=golang
 *
 * [589] N-ary Tree Preorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	res := []int{}

	var 	travel func(node *Node) 

	travel = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)

		for _, c := range node.Children {
			travel(c)
		}
	}
	travel(root)
	return res
    
}
// @lc code=end

