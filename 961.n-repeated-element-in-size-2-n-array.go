/*
 * @lc app=leetcode id=961 lang=golang
 *
 * [961] N-Repeated Element in Size 2N Array
 */

// @lc code=start
func repeatedNTimes(nums []int) int {
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			return n
		}
		seen[n] = true
	}
	return -1
}

// @lc code=end

