/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 */

// @lc code=start
func reverseStr(s string, k int) string {
	c := []byte(s)
	n := len(c)

	for i := 0; i < n; i += 2 * k {
		l, r := i, i+k-1
		if r >= n {
			r = n - 1
		}
		for l < r {
			c[l], c[r] = c[r], c[l]
			l++
			r--
		}
	}
	return string(c)
}

// @lc code=end

