/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	first := math.MaxInt
	second := math.MaxInt

	for _, x := range nums {
		if x <= first {
			first = x
		} else if x <= second {
			second = x
		} else {
			return true
		}
	}

	return false

}
// @lc code=end

