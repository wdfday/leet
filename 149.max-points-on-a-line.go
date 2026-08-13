/*
 * @lc app=leetcode id=149 lang=golang
 *
 * [149] Max Points on a Line
 */

// @lc code=start
func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	res := 1
	for i := 0; i < n; i++ {
		slope := make(map[[2]int]int)
		curMax := 0

		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			g := gcd(dx, dy)
			dx /= g
			dy /= g
			
			if dx < 0 || (dx == 0 && dy < 0) {
				dx, dy = -dx, -dy
			}

			key := [2]int{dx, dy}
			slope[key]++
			if slope[key] > curMax {
				curMax = slope[key]
			}
		}

		total := curMax + 1
		if total > res {
			res = total
		}
	}

	return res
}


func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	if a < 0 {
		return -a
	}
	return a
}
// @lc code=end

