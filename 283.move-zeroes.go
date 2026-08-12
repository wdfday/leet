/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	var l int

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			l = i
			break
		}
	}

	for i := l; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		} else {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}

	}

}

// @lc code=end

