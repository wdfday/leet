/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 */

// @lc code=start
func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, p := range []int{2, 3, 5} {
		for n%p == 0 {
			n /= p
		}
	}
	return n == 1
}

// @lc code=end

