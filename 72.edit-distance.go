/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m + 1)
	for i := range m + 1 {
		dp[i] = make([]int, n + 1)
	}

	for i := range m + 1 {
		for j := range n + 1 {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
// @lc code=end

