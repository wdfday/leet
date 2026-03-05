/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {

	var res []string
	var backtrack func(string, int, int)
	backtrack = func(curr string, open int, close int) {
		if len(curr) == 2*n {
			res = append(res, curr)
			return
		}
		if open < n {
			backtrack(curr+"(", open+1, close)
		}
		if close < open {
			backtrack(curr+")", open, close+1)
		}
	}
	backtrack("", 0, 0)
	return res
}

// @lc code=end

