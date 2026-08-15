/*
 * @lc app=leetcode id=3674 lang=golang
 *
 * [3674] Minimum Operations to Equalize Array
 */

// @lc code=start
func minOperations(nums []int) int {
    k := nums[0]
	for _, v := range nums {
		if v != k {
			return 1
		}
	}
	return 0
}
// @lc code=end

