/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
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

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil {
        return false
    }

    if subRoot == nil {
        return true
    }

    if IsSameTree(root, subRoot) {
        return true
    }

    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func IsSameTree(root *TreeNode, subRoot *TreeNode) bool {
    if root == nil && subRoot == nil {
        return true
    }

    if root == nil || subRoot == nil {
        return false
    }

    if root.Val != subRoot.Val {
        return false
    }

    return IsSameTree(root.Left, subRoot.Left) && IsSameTree(root.Right, subRoot.Right)
}


// @lc code=end


func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var preOrder func(node *TreeNode) string
	preOrder = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		// dùng dấu phân cách và marker để tránh false positive kiểu "12" chứa "2"
		return "^" + strconv.Itoa(node.Val) + " " + preOrder(node.Left) + " " + preOrder(node.Right)
	}

	tree := preOrder(root)
	sub := preOrder(subRoot)

	return strings.Contains(tree, sub)
}


func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	var preOrder func(node *TreeNode) []int 
	preOrder = func(node *TreeNode) []int {
		if node == nil {
			return []int{10001}
		} 
		t := []int{node.Val}
		t = append(t, preOrder(node.Left)...)
		t = append(t, preOrder(node.Right)...)
		return t
	}

	tree := preOrder(root)
	sub := preOrder(subRoot)

	return containsSeq(tree, sub)
}

// containsSeq kiểm tra sub có xuất hiện liên tiếp trong tree không
func containsSeq(tree, sub []int) bool {
	n, m := len(tree), len(sub)
	if m > n {
		return false
	}
	for i := 0; i+m <= n; i++ {
		match := true
		for j := 0; j < m; j++ {
			if tree[i+j] != sub[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}