/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	count := 1
	k := nums[0]
	res := len(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == k {
			count++
			if count == 3 {
				count--
				res--
				nums = append(nums[:i], nums[i+1:]...)
				i--
			}
		} else {
			k = nums[i]
			count = 1
		}
	}
	return res
}

// @lc code=end

