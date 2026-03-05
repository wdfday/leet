/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {

	left, right := 0, len(height)-1

	maxArea := 0
	for left < right {

		if height[left] < height[right] {
			area := height[left] * (right - left)
			if area > maxArea {
				maxArea = area
			}
			left++
		} else {
			area := height[right] * (right - left)
			if area > maxArea {
				maxArea = area
			}
			right--
		}

	}
	return maxArea

}

// @lc code=end
