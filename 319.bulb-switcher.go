/*
 * @lc app=leetcode id=319 lang=golang
 *
 * [319] Bulb Switcher
 */

// @lc code=start
func bulbSwitch(n int) int {
	r := int(math.Sqrt(float64(n)))
	for r*r > n {
		r--
	}
	for (r+1)*(r+1) <= n {
		r++
	}
	return r
}

// @lc code=end

