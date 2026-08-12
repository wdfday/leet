/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	l, r := 0, n-1
	for l < r {
		mid := l + (r-l)/2
		if citations[mid] >= n-mid {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if citations[l] >= n-l {
		return n - l
	}
	return 0
}

// @lc code=end

