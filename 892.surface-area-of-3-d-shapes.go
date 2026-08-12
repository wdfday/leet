/*
 * @lc app=leetcode id=892 lang=golang
 *
 * [892] Surface Area of 3D Shapes
 */

// @lc code=start
func surfaceArea(grid [][]int) int {
	n := len(grid)
	res := 0

	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				// top + bottom + 4 sides per cube stack base
				res += 2 + grid[i][j]*4
			}

			for k := 0; k < 4; k++ {
				ni, nj := i+dx[k], j+dy[k]
				if ni >= 0 && ni < n && nj >= 0 && nj < n {
					overlap := min(grid[i][j], grid[ni][nj])
					res -= overlap
				}
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

