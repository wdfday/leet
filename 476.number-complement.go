/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 */

// @lc code=start
func findComplement(num int) int {
	k := 1
	for num >= k {
		k *= 2
	}
	return k - num - 1

}

// @lc code=end

