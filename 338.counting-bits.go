/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
func countBits(n int) []int {

	res := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		res = append(res, countBit(i))
	}
	return res
}

func countBit(n int) int {
	res := 0

	for n > 0 {
		res += n & 1
		n >>= 1
	}
	return res
}

// @lc code=end

