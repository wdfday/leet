/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	m, n := len(v1), len(v2)
	k := max(m, n)

	for i := 0; i < k; i++ {
		a, b := 0, 0
		if i < m {
			a, _ = strconv.Atoi(v1[i])
		}
		if i < n {
			b, _ = strconv.Atoi(v2[i])
		}

		if a == b {
			continue
		} else if a > b {
			return 1
		} else {
			return -1
		}
	}

	return 0
}

// @lc code=end

