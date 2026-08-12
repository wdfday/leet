/*
 * @lc app=leetcode id=1047 lang=golang
 *
 * [1047] Remove All Adjacent Duplicates In String
 */

// @lc code=start
func removeDuplicates(s string) string {
	// Use a byte-stack to remove adjacent duplicates efficiently.
	res := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		b := s[i]
		if len(res) > 0 && res[len(res)-1] == b {
			res = res[:len(res)-1]
		} else {
			res = append(res, b)
		}
	}
	return string(res)
}
// @lc code=end

