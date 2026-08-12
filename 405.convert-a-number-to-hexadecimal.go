/*
 * @lc app=leetcode id=405 lang=golang
 *
 * [405] Convert a Number to Hexadecimal
 */

// @lc code=start
func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	hexChars := "0123456789abcdef"
	res := ""
	n := uint32(num) // ép về 32-bit unsigned

	for n != 0 {
		digit := n & 0xf
		res = string(hexChars[digit]) + res
		n >>= 4
	}

	return res
}

// @lc code=end

