/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func rob(nums []int) int {

	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	dp2 := make([]int, len(nums)+1)
	dp2[1] = nums[1]
	res := dp[0]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		dp2[i] = max(dp2[i-1], dp2[i-2]+nums[i])
		res := max(res, dp2[i])
		if i != len(nums)-1 {
			res = max(res, dp[i])
		}
	}

	return max(dp[len(nums)-2], dp2[len(nums)-1])
}

// @lc code=end

