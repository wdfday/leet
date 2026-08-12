/*
 * @lc app=leetcode id=3783 lang=golang
 *
 * [3783] Mirror Distance of an Integer
 */

// @lc code=start
func mirrorDistance(n int) int {
	mr := 0
	clone := n
	for clone > 0 {
		mr = mr*10 + clone%10
		clone /= 10

	}
	return abs(n - mr)
}

func abs(n int) int {
	if n < 0 {
		return -n

	}
	return n
}

// @lc code=end

