/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)

	maxSide := 0

	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min(
						dp[i-1][j],
						dp[i][j-1],
						dp[i-1][j-1],
					) + 1
				}
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}

	return maxSide * maxSide
}

// @lc code=end

