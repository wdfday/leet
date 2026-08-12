/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {

	n := len(nums)

	left := 0
	r := 0
	sum := 0
	result := math.MaxInt32

	for r < n {
		sum += nums[r]

		for sum >= target {
			result = min(result, r-left+1)
			sum -= nums[left]
			left++
		}

		r++
	}

	if result == math.MaxInt32 {
		return 0
	}

	return result

}

// @lc code=end

