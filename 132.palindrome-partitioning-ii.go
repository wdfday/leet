/*
 * @lc app=leetcode id=132 lang=golang
 *
 * [132] Palindrome Partitioning II
 */

// @lc code=start
func minCut(s string) int {
	m := len(s)

	dp := make([][]bool, m)

	for i := range m {
		dp[i] = make([]bool, m)
		dp[i][i] = true
	}

	for length := 2; length <= m; length ++ {
		for i := 0; i + length - 1 < m; i ++ {
			j := i + length - 1
			if s[i] == s[j] && (dp[i+1][j-1] || length <= 2) {
				dp[i][j] = true
			}
		}
	}

	cut := make([]int, m+1)
	for i := 0; i <= m; i++ {
		cut[i] = math.MaxInt
	}
	cut[0] = -1 
	for j := 0; j < m; j++ {
		for i := 0; i <= j; i++ {
			if dp[i][j] {
				cut[j+1] = min(cut[j+1], cut[i]+1)
			}
		}
	}

	return cut[m]
}
// @lc code=end

