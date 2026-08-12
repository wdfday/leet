/*
 * @lc app=leetcode id=3774 lang=golang
 *
 * [3774] Absolute Difference Between Maximum and Minimum K Elements
 */

// @lc code=start
func absDifference(nums []int, k int) int {
	sort.Ints(nums)
	l, s := 0, 0
	for i := 0; i < k; i++ {
		l += nums[len(nums)-1-i]
		s += nums[i]
	}
	return l - s
}

// @lc code=end

