/*
 * @lc app=leetcode id=365 lang=golang
 *
 * [365] Water and Jug Problem
 */

// @lc code=start
func canMeasureWater(x int, y int, target int) bool {
	if target > x+y {
		return false
	}
	if x == 0 || y == 0 {
		return target == 0 || target == x || target == y
	}
	return target%gcd(x, y) == 0
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// @lc code=end

func canMeasureWater(x, y, target int) bool {
	if target > x+y {
		return false
	}
	visited := make(map[[2]int]bool)
	queue := [][2]int{{0, 0}}
	visited[[2]int{0, 0}] = true

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		a, b := cur[0], cur[1]

		if a+b == target {
			return true // dừng sớm
		}

		for _, next := range neighbors(a, b, x, y) {
			if !visited[next] {
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	return false // duyệt hết state space, không tìm thấy
}

