/*
 * @lc app=leetcode id=297 lang=golang
 *
 * [297] Serialize and Deserialize Binary Tree
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
type Codec struct {
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	res := []string{}

	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			res = append(res, "#")
			return
		}
		res = append(res, strconv.Itoa(node.Val))
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return strings.Join(res, ",")

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	idx := 0 

	var build func() *TreeNode
	build = func() *TreeNode {
		if idx >= len(vals) {
			return nil
		}
		v := vals[idx]
		idx++
		if v == "#" {
			return nil
		}
		num, _ := strconv.Atoi(v)
		node := &TreeNode{Val: num}
		node.Left = build()
		node.Right = build()
		return node
	}

	return build()

    
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

