/*
 * @lc app=leetcode id=1009 lang=golang
 *
 * [1009] Complement of Base 10 Integer
 */

// @lc code=start
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	k := 1
	for k <= n {
		k *= 2
	}

	return k - n - 1

}

// @lc code=end

