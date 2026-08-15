/*
 * @lc app=leetcode id=3643 lang=golang
 *
 * [3643] Flip Square Submatrix Vertically
 */

// @lc code=start
func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for i := 0; i < k; i++ {
		for j := 0; j < k/2; j++ {
			col := y + i
			grid[x+j][col], grid[x+k-j-1][col] = grid[x+k-j-1][col], grid[x+j][col]
		}
	}
	return grid
}
// @lc code=end

