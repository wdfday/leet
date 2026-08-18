/*
 * @lc app=leetcode id=97 lang=golang
 *
 * [97] Interleaving String
 */

// @lc code=start
func isInterleave(s1 string, s2 string, s3 string) bool {
	m := len(s1)
	n := len(s2)

	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			k := i + j - 1
            
			f1 := dp[i-1][j] && s1[i-1] == s3[k]
			f2 := dp[i][j-1] && s2[j-1] == s3[k]

			dp[i][j] = f1 || f2
		}
	}

	return dp[m][n]
}
// @lc code=end

func isInterleave(s1 string, s2 string, s3 string) bool {
    m, n := len(s1), len(s2)
    if m+n != len(s3) {
        return false
    }
    dp := make([]bool, n+1)
    dp[0] = true
    for j := 1; j <= n; j++ {
        dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
    }
    for i := 1; i <= m; i++ {
        dp[0] = dp[0] && s1[i-1] == s3[i-1]
        for j := 1; j <= n; j++ {
            dp[j] = (dp[j] && s1[i-1] == s3[i+j-1]) ||
                    (dp[j-1] && s2[j-1] == s3[i+j-1])
        }
    }
    return dp[n]
}