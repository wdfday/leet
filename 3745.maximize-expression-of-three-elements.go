/*
 * @lc app=leetcode id=3745 lang=golang
 *
 * [3745] Maximize Expression of Three Elements
 */

// @lc code=start
func maximizeExpressionOfThree(nums []int) int {
    slices.Sort(nums)

	return nums[len(nums) - 1] +nums[len(nums) - 2] - nums[0]
}
// @lc code=end

