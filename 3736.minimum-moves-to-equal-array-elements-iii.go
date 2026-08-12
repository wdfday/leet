/*
 * @lc app=leetcode id=3736 lang=golang
 *
 * [3736] Minimum Moves to Equal Array Elements III
 */

// @lc code=start
func minMoves(nums []int) int {
	max := nums[0]
	res := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			res += i * (nums[i] - max)
			max = nums[i]
		} else {
			res += max - nums[i]
		}
	}
	return res
}

// @lc code=end

