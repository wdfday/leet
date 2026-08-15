/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */

// @lc code=start
func partition(s string) [][]string {
    m := len(s)
	dp := make([][]bool, m)
	res := [][]string{}
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

	var backtrack func(j int, par []string)
	backtrack = func(j int, par []string) {
		if j == m  {
		    res = append(res, append([]string(nil), par...))
			return
		}
		for k := j; k < m; k++ {
			if dp[j][k] {
				sub := s[j:k+1]
				backtrack(k+1, append(par, sub))
			}
		}
	}

	backtrack(0, []string{})
	return res
}
// @lc code=end

