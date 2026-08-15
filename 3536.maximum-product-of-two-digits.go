/*
 * @lc app=leetcode id=3536 lang=golang
 *
 * [3536] Maximum Product of Two Digits
 */

// @lc code=start
func maxProduct(n int) int {
	m1 := n%10
	m := 0
	n /= 10
	for n > 0 {
		t1 := n % 10
		n/=10

		m = max(m, t1 * m1)
		m1 = max(t1, m1)
	}
	return m
}
// @lc code=end

