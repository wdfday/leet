/*
 * @lc app=leetcode id=3028 lang=golang
 *
 * [3028] Ant on the Boundary
 */

// @lc code=start
func returnToBoundaryCount(nums []int) int {
	k := 0
	res := 0
	for _, v := range nums {
		k += v
		if k == 0 {
			res++
		}
	}
    return res
}
// @lc code=end

