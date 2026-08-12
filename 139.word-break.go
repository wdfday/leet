/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]struct{}, len(wordDict))
	for _, w := range wordDict {
		dict[w] = struct{}{}
	}

	dp := map[int]bool{-1: true}

	for i := range s {
		for k := range dp {
			if _, ok := dict[s[k+1:i+1]]; ok {
				dp[i] = true
				break // i đã reachable thì không cần thử tiếp
			}
		}
	}

	return dp[len(s)-1]

}

// @lc code=end

func wordBreak(s string, wordDict []string) bool {
	n := len(s)

	dict := make(map[string]struct{}, len(wordDict))
	maxLen := 0

	for _, w := range wordDict {
		dict[w] = struct{}{}
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		start := max(0, i-maxLen)

		for j := start; j < i; j++ {
			if !dp[j] {
				continue
			}

			if _, ok := dict[s[j:i]]; ok {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}