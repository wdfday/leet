/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
func majorityElement(nums []int) int {
	count := 1
	k := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == k {
			count++
		} else {
			count--
		}

		if count == 0 {
			k = nums[i]
			count = 1
		}
	}

	return k
}

// @lc code=end

