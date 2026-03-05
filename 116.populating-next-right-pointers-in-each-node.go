/*
 * @lc app=leetcode id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}
// 	var dfs func(left, right *Node)
// 	dfs = func(left, right *Node) {
// 		if left == nil || right == nil {
// 			return
// 		}

// 		left.Next = right
// 		dfs(left.Left, left.Right)
// 		dfs(right.Left, right.Right)
// 		dfs(left.Right, right.Left)
// 	}

// 	dfs(root.Left, root.Right)
// 	return root
// }

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		var prev *Node

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if prev != nil {
				prev.Next = node
			}
			prev = node

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}

/*
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if i < size-1 {
				node.Next = queue[0]
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return root
}
*/

// @lc code=end
