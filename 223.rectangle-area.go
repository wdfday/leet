/*
 * @lc app=leetcode id=223 lang=golang
 *
 * [223] Rectangle Area
 */

// @lc code=start
func computeArea(ax1, ay1, ax2, ay2 int,
	bx1, by1, bx2, by2 int) int {

	areaA := (ax2 - ax1) * (ay2 - ay1)
	areaB := (bx2 - bx1) * (by2 - by1)

	overlapX := min(ax2, bx2) - max(ax1, bx1)
	overlapY := min(ay2, by2) - max(ay1, by1)

	overlap := 0
	if overlapX > 0 && overlapY > 0 {
		overlap = overlapX * overlapY
	}

	return areaA + areaB - overlap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

