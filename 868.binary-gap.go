/*
 * @lc app=leetcode id=868 lang=golang
 *
 * [868] Binary Gap
 */

// @lc code=start
func binaryGap(n int) int {
	k := -1
	res := 0
	for n > 0 {
		isB := n & 1
		n /= 2
		if isB == 1 {
			if k == -1 {
				k = 1
			} else {
				res = max(k, res)
				k = 1
			}
		} else {
			if k == -1 {
				continue
			} else {
				k++
			}
		}

	}
	return res
}

// @lc code=end

