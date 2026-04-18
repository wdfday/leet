/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {

	l, r := 0, len(nums)-1
	i := 0

	for i <= r {
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			i++
			l++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else {
			i++
		}
	}
}

// @lc code=end

