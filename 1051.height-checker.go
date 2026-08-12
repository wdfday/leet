/*
 * @lc app=leetcode id=1051 lang=golang
 *
 * [1051] Height Checker
 */

// @lc code=start
func heightChecker(heights []int) int {
    expect := make([]int, len(heights))
	copy(expect, heights)
	slices.Sort(expect)
	res := 0
	for i, v := range heights {
		if v != expect[i] {
			res++
		}
	}
	return res
}
// @lc code=end

