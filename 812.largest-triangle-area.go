/*
 * @lc app=leetcode id=812 lang=golang
 *
 * [812] Largest Triangle Area
 */

// @lc code=start
func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	max := 0.0

	area := func(a, b, c []int) float64 {
		return math.Abs(float64(
			a[0]*(b[1]-c[1])+
				b[0]*(c[1]-a[1])+
				c[0]*(a[1]-b[1]),
		)) / 2
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if s := area(points[i], points[j], points[k]); s > max {
					max = s
				}
			}
		}
	}

	return max
}

// @lc code=end

