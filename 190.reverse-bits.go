/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(n int) int {
	u := uint32(n)
	var r uint32
	for i := 0; i < 32; i++ {
		r = (r << 1) | (u & 1)
		u >>= 1
	}
	return int(r)
}

// @lc code=end

