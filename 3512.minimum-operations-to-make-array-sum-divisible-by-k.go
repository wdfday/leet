/*
 * @lc app=leetcode id=3512 lang=golang
 *
 * [3512] Minimum Operations to Make Array Sum Divisible by K
 */

// @lc code=start
func minOperations(nums []int, k int) int {
	sum := 0

	for _, v := range nums {
		sum += v 
	}
	return sum % k
    
}
// @lc code=end

