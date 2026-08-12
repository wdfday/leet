/*
 * @lc app=leetcode id=3798 lang=golang
 *
 * [3798] Largest Even Number
 */

// @lc code=start
func largestEven(s string) string {
	if len(s) == 1 {
		if s[0] == '2' {
			return s
		}
		return ""
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '2' {
			return s[:i+1]
		}
	}
	return ""
}

// @lc code=end

