/*
 * @lc app=leetcode id=3712 lang=golang
 *
 * [3712] Sum of Elements With Frequency Divisible by K
 */

// @lc code=start
func sumDivisibleByK(nums []int, k int) int {
	frq := make(map[int]int)
	res := 0
	for _,v := range nums {
		frq[v]++
	}
	for i, v := range frq {
		if v % k == 0 {
			res += i * v
		}
	}
	return res  
}
// @lc code=end

