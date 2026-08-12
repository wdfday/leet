/*
 * @lc app=leetcode id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	l, r := 0, num
	for l <= r {
		m := l + (r-l)/2
		sq := m * m
		if sq == num {
			return true
		}
		if sq < num {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return false

}

// @lc code=end

