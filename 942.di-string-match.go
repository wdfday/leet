/*
 * @lc app=leetcode id=942 lang=golang
 *
 * [942] DI String Match
 */

// @lc code=start
func diStringMatch(s string) []int {
	n := len(s)
	ans := make([]int, 0, n+1)

	low, high := 0, n

	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			ans = append(ans, low)
			low++
		} else {
			ans = append(ans, high)
			high--
		}
	}

	ans = append(ans, low)

	return ans
}

// @lc code=end

