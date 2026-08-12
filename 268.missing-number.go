/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {

	sum := len(nums)
	for i := 0; i < len(nums); i++ {
		sum = sum + i - nums[i]
	}

	return sum

}

// @lc code=end

