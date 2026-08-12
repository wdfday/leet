/*
 * @lc app=leetcode id=944 lang=golang
 *
 * [944] Delete Columns to Make Sorted
 */

// @lc code=start
func minDeletionSize(strs []string) int {
	m, n := len(strs), len(strs[0])
	res := 0

	for j := 0; j < n; j++ {
		for i := 1; i < m; i++ {
			if strs[i][j] < strs[i-1][j] {
				res++
				break
			}
		}
	}
	return res
}

// @lc code=end

