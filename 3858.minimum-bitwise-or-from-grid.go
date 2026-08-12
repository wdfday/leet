/*
 * @lc app=leetcode id=3858 lang=golang
 *
 * [3858] Minimum Bitwise OR From Grid
 */

// @lc code=start
func minimumOR(grid [][]int) int {
	m := len(grid)
	ans := 0
	for b := 17; b >= 0; b-- {
		mask := ^((1 << b) - 1) // giữ nguyên bit từ vị trí b trở lên, xóa các bit thấp hơn

		ok := true
		for i := 0; i < m && ok; i++ {
			found := false
			for _, x := range grid[i] {
				if (x & mask) == ans {
					found = true
					break
				}
			}
			if !found {
				ok = false
			}
		}

		if !ok {
			ans |= (1 << b)
		}
	}
	return ans

}

// @lc code=end
