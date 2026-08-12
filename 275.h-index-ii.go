/*
 * @lc app=leetcode id=275 lang=golang
 *
 * [275] H-Index II
 */

// @lc code=start
func hIndex(citations []int) int {
	n := len(citations)
	l, r := 0, n

	for l < r {
		mid := (l + r) / 2

		if citations[mid] >= n-mid {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return n - l
}

// @lc code=end

