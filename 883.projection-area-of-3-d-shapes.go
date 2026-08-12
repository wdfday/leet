/*
 * @lc app=leetcode id=883 lang=golang
 *
 * [883] Projection Area of 3D Shapes
 */

// @lc code=start
func projectionArea(grid [][]int) int {
	n := len(grid)

	top := 0
	rowMax := 0
	colMax := make([]int, n)

	for i := 0; i < n; i++ {
		rMax := 0
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				top++
			}
			if grid[i][j] > rMax {
				rMax = grid[i][j]
			}
			if grid[i][j] > colMax[j] {
				colMax[j] = grid[i][j]
			}
		}
		rowMax += rMax
	}

	colSum := 0
	for j := 0; j < n; j++ {
		colSum += colMax[j]
	}

	return top + rowMax + colSum
}

// @lc code=end

