/*
 * @lc app=leetcode id=3668 lang=golang
 *
 * [3668] Restore Finishing Order
 */

// @lc code=start
func recoverOrder(order []int, friends []int) []int {
	res := []int{}

	m := make(map[int]struct{})
	for _, f := range friends {
		m[f] = struct{}{}
	}

	for _, v := range order {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}
    return res
}
// @lc code=end

