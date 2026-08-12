/*
 * @lc app=leetcode id=3790 lang=golang
 *
 * [3790] Smallest All-Ones Multiple
 */

// @lc code=start
func minAllOneMultiple(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	v := 0
	for i := range k {
		v = (v*10 + 1) % k
		if v == 0 {
			return i + 1
		}
	}
	return -1
}

// @lc code=end

