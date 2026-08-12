/*
 * @lc app=leetcode id=628 lang=golang
 *
 * [628] Maximum Product of Three Numbers
 */

// @lc code=start
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	candidate1 := nums[0] * nums[1] * nums[n-1]
	candidate2 := nums[n-1] * nums[n-2] * nums[n-3]

	return max(candidate1, candidate2)
}

// @lc code=end

