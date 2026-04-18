/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	can := nums[0]
	n := len(nums) - 1
	for i := 0; i <= can && i <= n; i++ {
		can = max(i+nums[i], can)
	}

	if can >= n {
		return true
	}

	return false

}

// @lc code=end

