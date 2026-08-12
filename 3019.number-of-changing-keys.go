/*
 * @lc app=leetcode id=3019 lang=golang
 *
 * [3019] Number of Changing Keys
 */

// @lc code=start
func countKeyChanges(s string) int {
	res := 0

	for i := 1; i < len(s); i++ {
		if s[i]|32 != s[i-1]|32 {
			res++
		}
	}

	return res

}
// @lc code=end

