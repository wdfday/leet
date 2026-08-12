/*
 * @lc app=leetcode id=3795 lang=golang
 *
 * [3795] Minimum Subarray Length With Distinct Sum At Least K
 */

// @lc code=start
func minLength(nums []int, k int) int {
	l := 0
	freq := make(map[int]int)
	sum := 0
	res := -1

	for r := 0; r < len(nums); r++ {
		if freq[nums[r]] == 0 {
			sum += nums[r]
		}
		freq[nums[r]]++

		for sum >= k {
			length := r - l + 1
			if res == -1 || length < res {
				res = length
			}
			freq[nums[l]]--
			if freq[nums[l]] == 0 {
				sum -= nums[l]
			}
			l++
		}
	}
	return res
}

// @lc code=end

