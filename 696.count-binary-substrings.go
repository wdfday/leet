/*
 * @lc app=leetcode id=696 lang=golang
 *
 * [696] Count Binary Substrings
 */

// @lc code=start
func countBinarySubstrings(s string) int {
	res := 0
	consect, consecf := 0, 0
	for i := 0; i < len(s); {
		if s[i] == '1' {
			for i < len(s) && s[i] == '1'   {
				consect++
				i++
			}
			res += min(consect, consecf)
			consecf = 0
		} else {
			for i < len(s) && s[i] == '0' {
				consecf++
				i++
			}
			res += min(consect, consecf)
			consect = 0
		}
	}
	return res
}

// @lc code=end

