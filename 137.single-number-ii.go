/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, x := range nums {
		ones = (ones ^ x) & ^twos
		twos = (twos ^ x) & ^ones
	}
	return ones

}

// @lc code=end

