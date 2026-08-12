/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */

// @lc code=start
// func isPowerOfFour(n int) bool {
// 	return n > 0 && (n&(n-1)) == 0 && (n&0x55555555) != 0
// }

func isPowerOfFour(n int) bool {
	for x := 1; x <= n; x *= 4 {
		if x == n {
			return true
		}
	}
	return false
}

// @lc code=end

