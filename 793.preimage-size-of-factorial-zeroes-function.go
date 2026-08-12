/*
 * @lc app=leetcode id=793 lang=golang
 *
 * [793] Preimage Size of Factorial Zeroes Function
 */

// @lc code=start
func preimageSizeFZF(k int) int {
	zeros := func(n int64) int64 {
		var count int64
		for n > 0 {
			n /= 5
			count += n
		}
		return count
	}

	lo, hi := int64(0), int64(5)*int64(k)+5
	for lo < hi {
		mid := lo + (hi-lo)/2
		if zeros(mid) < int64(k) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	if zeros(lo) == int64(k) {
		return 5
	}
	return 0
}

// @lc code=end

