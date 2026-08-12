/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	m := sum

	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		m = max(sum, m)
	}

	return float64(m) / float64(k)
}

// @lc code=end

