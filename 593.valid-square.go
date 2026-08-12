/*
 * @lc app=leetcode id=593 lang=golang
 *
 * [593] Valid Square
 */

// @lc code=start
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {

	dist := func(a, b []int) int {
		dx := a[0] - b[0]
		dy := a[1] - b[1]
		return dx*dx + dy*dy
	}

	points := [][]int{p1, p2, p3, p4}
	distances := make(map[int]int)

	for i := 0; i < 4; i++ {
		for j := i + 1; j < 4; j++ {
			d := dist(points[i], points[j])
			if d == 0 {
				return false
			}
			distances[d]++
		}
	}

	if len(distances) != 2 {
		return false
	}

	var side, diagonal int
	for d, count := range distances {
		if count == 4 {
			side = d
		} else if count == 2 {
			diagonal = d
		} else {
			return false
		}
	}

	return diagonal == 2*side
}

// @lc code=end

