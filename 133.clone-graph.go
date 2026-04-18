/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	return clone(node, make(map[int]*Node))
}

func clone(node *Node, visited map[int]*Node) *Node {
	if node == nil {
		return nil
	}

	if v, ok := visited[node.Val]; ok {
		return v
	}

	newNode := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	visited[node.Val] = newNode

	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, clone(neighbor, visited))
	}

	return newNode
}

// @lc code=end

