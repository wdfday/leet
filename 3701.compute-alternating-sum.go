/*
 * @lc app=leetcode id=3701 lang=golang
 *
 * [3701] Compute Alternating Sum
 */

// @lc code=start
func alternatingSum(nums []int) int {
    res := 0
	for i, v := range nums {
		if i % 2 == 0 {
			res += v
		} else {
			res -= v
		}
	}
	return res
}
// @lc code=end

