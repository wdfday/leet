/*
 * @lc app=leetcode id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	consec := 1
	res := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			consec++
			res = max(consec, res)

		} else {
			consec = 1
		}
	}
	return res

}

// @lc code=end

