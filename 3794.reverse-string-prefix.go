/*
 * @lc app=leetcode id=3794 lang=golang
 *
 * [3794] Reverse String Prefix
 */

// @lc code=start
func reversePrefix(s string, k int) string {
	c := []byte(s)
	l, r := 0, k-1
	for l < r {
		c[l], c[r] = c[r], c[l]
		l++
		r--
	}

	return string(c)

}

// @lc code=end

