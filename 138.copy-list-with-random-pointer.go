/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// original node → copied node
	nodeMap := make(map[*Node]*Node)

	cur := head
	for cur != nil {
		nodeMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		if cur.Next != nil {
			nodeMap[cur].Next = nodeMap[cur.Next]
		}
		if cur.Random != nil {
			nodeMap[cur].Random = nodeMap[cur.Random]
		}
		cur = cur.Next
	}

	return nodeMap[head]
}

// @lc code=end

