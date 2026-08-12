/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	z := x ^ y
	res := 0
	for z > 0 {
		if z%2 == 1 {
			res++
		}
		z /= 2
	}
	return res
}

// @lc code=end

