/*
 * @lc app=leetcode id=3688 lang=golang
 *
 * [3688] Bitwise OR of Even Numbers in an Array
 */

// @lc code=start
func evenNumberBitwiseORs(nums []int) int {
	res := 0
	for _, num := range nums {
		if num % 2 ==0 {
			res = res | num
		}
	}
    return res
}
// @lc code=end

