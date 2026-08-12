/*
 * @lc app=leetcode id=492 lang=golang
 *
 * [492] Construct the Rectangle
 */

// @lc code=start
func constructRectangle(area int) []int {

	a := int(math.Sqrt(float64(area)))

	for a > 1 {
		b := area / a
		if a*b == area {
			return []int{b, a}
		}
		a--

	}

	return []int{area, 1}
}

// @lc code=end

