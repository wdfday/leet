/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {

	res, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		cur = max(cur+nums[i], nums[i])
		res = max(cur, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

