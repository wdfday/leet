/*
 * @lc app=leetcode id=1745 lang=golang
 *
 * [1745] Palindrome Partitioning IV
 */

// @lc code=start
func checkPartitioning(s string) bool {
	m := len(s)
	dp := make([][]bool, m)
	for i := range m {
		dp[i] = make([]bool, m)
		dp[i][i] = true
	}

	for i := range m - 1 {

		j, k := i-1, i+1
		for j >= 0 && k < m {
			if s[j] == s[k] && dp[j+1][k-1] {
				dp[j][k] = true
				j--
				k++
			} else {
				break
			}
		}
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
		j, k = i-1, i+2
		for j >= 0 && k < m {
			if s[j] == s[k] && dp[j+1][k-1] {
				dp[j][k] = true
				j--
				k++
			} else {
				break
			}
		}
	}

	for i := range m - 2 {
		for j := i + 1; j < m-1; j++ {
			if dp[0][i] && dp[i+1][j] && dp[j+1][m-1] {
				return true
			}
		}
	}
	return false
}

// @lc code=end

