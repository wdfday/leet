/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j := 0, 0
	lg, ls := len(g), len(s)
	res := 0

	for j < ls && i < lg {
		if g[i] <= s[j] {
			i++
			j++
			res++
		} else {
			j++
		}

	}
	return res
}

// @lc code=end

