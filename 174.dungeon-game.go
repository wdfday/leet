/*
 * @lc app=leetcode id=174 lang=golang
 *
 * [174] Dungeon Game
 */

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	const INF = int(1e9)

	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	dp[m][n-1] = 1
	dp[m-1][n] = 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			need := min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			dp[i][j] = max(1, need)
		}
	}

	return dp[0][0]
}
// @lc code=end

