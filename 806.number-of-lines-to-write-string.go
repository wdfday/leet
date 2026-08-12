/*
 * @lc app=leetcode id=806 lang=golang
 *
 * [806] Number of Lines To Write String
 */

// @lc code=start
func numberOfLines(widths []int, s string) []int {
	line := 1
	px := 0
	for _, b := range s {
		if widths[int(b-'a')] + px > 100 {
			line++
			px = widths[int(b-'a')]
		} else {
			px += widths[int(b-'a')]
		}
	}
	return []int{line, px}
    
}
// @lc code=end

