/*
 * @lc app=leetcode id=896 lang=golang
 *
 * [896] Monotonic Array
 */

// @lc code=start
func isMonotonic(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	dir := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if dir == 0 {
				dir = 1
			} else if dir == -1 {
				return false
			}
		} else if nums[i] < nums[i-1] {
			if dir == 0 {
				dir = -1
			} else if dir == 1 {
				return false
			}
		}
	}
	return true
}

// @lc code=end

