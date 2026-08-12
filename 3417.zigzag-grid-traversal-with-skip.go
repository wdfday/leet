/*
 * @lc app=leetcode id=3417 lang=golang
 *
 * [3417] Zigzag Grid Traversal With Skip
 */

// @lc code=start
func zigzagTraversal(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	res := make([]int, 0)
	idx := 0
	for i := 0; i < m; i++ {
		if i%2 == 0 {
			for j := 0; j < n; j++ {
				if idx%2 == 0 {
					res = append(res, grid[i][j])
				}
				idx++
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				if idx%2 == 0 {
					res = append(res, grid[i][j])
				}
				idx++
			}
		}
	}
	return res
}

// @lc code=end

