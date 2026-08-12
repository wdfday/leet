/*
 * @lc app=leetcode id=766 lang=golang
 *
 * [766] Toeplitz Matrix
 */

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		if i == 0 {
			for j := 0; j < n; j++ {
				x, y := i, j
				for {
					if y == n || x == m {
						break
					}
					if matrix[x][y] != matrix[i][j] {
						return false
					}
					x++
					y++
				}
			}
		} else {
			x, y := i, 0
			for {
				if y == n || x == m {
					break
				}
				if matrix[x][y] != matrix[i][0] {
					return false
				}
				x++
				y++
			}
		}
	}
	return true
}

// @lc code=end

