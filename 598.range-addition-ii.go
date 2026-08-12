/*
 * @lc app=leetcode id=598 lang=golang
 *
 * [598] Range Addition II
 */

// @lc code=start
func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}
	minx, miny := ops[0][0], ops[0][1]
	for _, op := range ops {
		if minx > op[0] {
			minx = op[0]
		}
		if miny > op[1] {
			miny = op[1]
		}
	}

	return minx * miny

}

// @lc code=end

