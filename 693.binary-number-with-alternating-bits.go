/*
 * @lc app=leetcode id=693 lang=golang
 *
 * [693] Binary Number with Alternating Bits
 */

// @lc code=start
func hasAlternatingBits(n int) bool {
	b := n % 2
	n /= 2
	for n > 0 {
		if n%2 == b {
			return false
		}
		b = n % 2
		n /= 2
	}
	return true
}

// @lc code=end

