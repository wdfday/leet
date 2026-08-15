/*
 * @lc app=leetcode id=301 lang=golang
 *
 * [301] Remove Invalid Parentheses
 */

// @lc code=start
func removeInvalidParentheses(s string) []string {
	res := []string{}

	var dfs func(string, int)
	dfs = func(s string, start int) {
		cnt := 0

		for i := start; i < len(s); i++ {
			switch s[i] {
			case '(':
				cnt++
			case ')':
				cnt--
			}

			// Prefix invalid: must delete a ')'
			if cnt < 0 {
				for j := start; j <= i; j++ {
					if s[j] != ')' {
						continue
					}

					// Avoid deleting consecutive ')' redundantly.
					if j > start && s[j-1] == ')' {
						continue
					}

					dfs(s[:j]+s[j+1:], i)
				}

				return
			}
		}

		// No excess ')'. Now deal with excess '('.
		if cnt > 0 {
			for i := len(s) - 1; i >= 0; i-- {
				if s[i] != '(' {
					continue
				}

				// Avoid deleting consecutive '(' redundantly.
				if i+1 < len(s) && s[i+1] == '(' {
					continue
				}

				dfs(s[:i]+s[i+1:], i)
			}

			return
		}

		res = append(res, s)
	}

	dfs(s, 0)
	return res
}
// @lc code=end

