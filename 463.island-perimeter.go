/*
 * @lc app=leetcode id=463 lang=golang
 *
 * [463] Island Perimeter
 */

// @lc code=start
func islandPerimeter(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// up
				if i == 0 || grid[i-1][j] == 0 {
					res++
				}
				// down
				if i == m-1 || grid[i+1][j] == 0 {
					res++
				}
				// left
				if j == 0 || grid[i][j-1] == 0 {
					res++
				}
				// right
				if j == n-1 || grid[i][j+1] == 0 {
					res++
				}
			}
		}
	}
	return res
}
// @lc code=end

