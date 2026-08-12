/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	slow, fast := n, next(n)
	for slow != fast && fast != 1 {
		slow = next(slow)
		fast = next(next(fast))
	}
	return fast == 1 || slow == 1
}

func next(n int) int {
	s := 0
	for n > 0 {
		d := n % 10
		s += d * d
		n /= 10
	}
	return s
}

// @lc code=end

