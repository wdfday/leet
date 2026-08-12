/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	m, n := len(s), len(t)
	if m < n {
		return ""
	}

	mapT := make(map[byte]int)

	for i := 0; i < n; i++ {
		mapT[t[i]]++
	}
	required := len(mapT) // số ký tự DISTINCT cần đủ
	formed := 0

	mapW := make(map[byte]int) // window hiện tại

	l, r := 0, 0
	res := make([]int, 2)
	res[1] = math.MaxInt32

	for r < m {

		// khi thêm s[r] vào window:
		mapW[s[r]]++
		if mapW[s[r]] == mapT[s[r]] { // vừa đủ số lượng ký tự này
			formed++
		}

		for formed == required {
			if r-l < res[1]-res[0] {
				res[0], res[1] = l, r
			}

			mapW[s[l]]--
			if mapW[s[l]] < mapT[s[l]] { // vừa đủ số lượng ký tự này
				formed--
			}
			l++
		}
		r++
	}
	if res[1] == math.MaxInt32 {
		return ""
	}
	return s[res[0] : res[1]+1]
}

// @lc code=end

