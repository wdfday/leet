/*
 * @lc app=leetcode id=590 lang=golang
 *
 * [590] N-ary Tree Postorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    res := []int{}

	var travel func(node *Node) 
	travel = func(node *Node) {
		if node == nil {
			return
		}

		for _, c := range node.Children {
			travel(c)
		}
		res = append(res, node.Val)

	}
	travel(root)
	return res

}
// @lc code=end

