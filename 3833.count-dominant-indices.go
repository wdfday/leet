/*
 * @lc app=leetcode id=3833 lang=golang
 *
 * [3833] Count Dominant Indices
 */

// @lc code=start
func dominantIndices(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	avg := float64(nums[n-1])
	res := 0

	for i := n - 2; i >= 0; i-- {
		if float64(nums[i]) > avg {
			res++
		}
		// avg của nums[i..n-1]
		avg = (avg*float64(n-i-1) + float64(nums[i])) / float64(n-i)
	}

	return res
}

// @lc code=end

