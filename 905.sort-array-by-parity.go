/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 */

// @lc code=start
func sortArrayByParity(nums []int) []int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		if nums[lo]%2 == 0 {
			lo++
		} else {
			nums[lo], nums[hi] = nums[hi], nums[lo]
			hi--
		}
	}
	return nums
}

// @lc code=end

