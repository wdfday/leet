/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 */

// @lc code=start
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	maxIndex := 0
	m := 1
	for i := range n {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] % nums[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > m {
			m = dp[i]
			maxIndex = i
		}
	}
	res := []int{nums[maxIndex]}
	for i := maxIndex - 1; i >= 0 ; i-- {
		if (nums[maxIndex] % nums[i] == 0) && (dp[maxIndex] == dp[i] + 1) {
			res = append([]int{nums[i]}, res...)
			maxIndex = i
		}
	}

	return res
}
// @lc code=end

