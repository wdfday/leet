/*
 * @lc app=leetcode id=3516 lang=golang
 *
 * [3516] Find Closest Person
 */

// @lc code=start
func findClosest(x int, y int, z int) int {
	x = abs(x - z)
	y = abs(y - z)
	if x > y {
		return 2
	} else if x < y {
		return 1
	}
	return 0
}

func abs(a int)int {
	if a < 0 {
		return -a 
	}
	return a
}

// @lc code=end

