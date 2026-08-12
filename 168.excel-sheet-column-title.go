/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
func convertToTitle(columnNumber int) string {
	if columnNumber <= 0 {
		return ""
	}
	res := make([]byte, 0)
	for columnNumber > 0 {
		columnNumber-- // shift to 0..25
		res = append(res, byte('A'+(columnNumber%26)))
		columnNumber /= 26
	}
	// reverse
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)

}

// @lc code=end

