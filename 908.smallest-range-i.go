/*
 * @lc app=leetcode id=908 lang=golang
 *
 * [908] Smallest Range I
 */

// @lc code=start
func smallestRangeI(nums []int, k int) int {
	if len(nums) == 1 {
		return 0
	}
	mn, mx := nums[0], nums[0]
	for _, n := range nums {
		if n < mn {
			mn = n
		}
		if n > mx {
			mx = n
		}
	}
	if mx-mn <= 2*k {
		return 0
	}
	return mx - mn - 2*k
}

// @lc code=end

