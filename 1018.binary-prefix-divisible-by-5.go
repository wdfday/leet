/*
 * @lc app=leetcode id=1018 lang=golang
 *
 * [1018] Binary Prefix Divisible By 5
 */

// @lc code=start
func prefixesDivBy5(nums []int) []bool {
	k := 0
	res := make([]bool, len(nums))
	for i, v := range nums {
		k = (k*2 + v) % 5
		res[i] = k == 0
	}
	return res
}
// @lc code=end

