/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	carry := 1
	i := len(digits) - 1
	for carry == 1 && i >= 0 {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10

		if i == 0 && carry == 1 {
			digits = append([]int{1}, digits...)
		}
		i--
	}

	return digits
}

// @lc code=end

