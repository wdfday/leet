/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {

	l, r := 0, len(height)-1
	res := (r - l) * min(height[l], height[r])
	var area int

	for l < r {
		if height[l] < height[r] {
			l++
			area = (r - l) * min(height[l], height[r])
			res = max(res, area)
		} else {
			r--
			area = (r - l) * min(height[l], height[r])
			res = max(res, area)

		}
	}
	return res
}

// @lc code=end
