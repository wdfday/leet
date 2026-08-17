/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */

// @lc code=start
func getSum(a int, b int) int {
    for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a

}
// @lc code=end

