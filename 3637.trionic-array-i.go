/*
 * @lc app=leetcode id=3637 lang=golang
 *
 * [3637] Trionic Array I
 */

// @lc code=start
func isTrionic(nums []int) bool {
	if nums[1] <= nums[0] {
		return false
	}
    m := 0 
	for i := 2; i < len(nums); i ++ {
		if nums[i] == nums[i-1] {
			return false
		}
		switch m {
		case 0:
			if nums[i] < nums[i-1] {
				m++
			} 
		case 1: 
			if nums[i] > nums[i-1] {
				m++
			}
		case 2: 
			if nums[i] < nums[i-1] {
				return false
			}
		}
	}
	return m == 2
}
// @lc code=end

