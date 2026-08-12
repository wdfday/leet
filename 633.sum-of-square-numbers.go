/*
 * @lc app=leetcode id=633 lang=golang
 *
 * [633] Sum of Square Numbers
 */

// @lc code=start
func judgeSquareSum(c int) bool {
	l, r := 0, int(math.Sqrt(float64(c)))

	for l <= r {
		sum := l*l + r*r
		if sum == c {
			return true
		} else if sum < c {
			l++
		} else {
			r--
		}
	}
	return false
}

// @lc code=end

