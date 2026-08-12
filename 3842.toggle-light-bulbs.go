/*
 * @lc app=leetcode id=3842 lang=golang
 *
 * [3842] Toggle Light Bulbs
 */

// @lc code=start
func toggleLightBulbs(bulbs []int) []int {
	m := make(map[int]bool)
	res := make([]int, 0)

	for i := 0; i < len(bulbs); i++ {
		m[bulbs[i]] = !m[bulbs[i]]
	}

	for k, v := range m {
		if v == true {
			res = append(res, k)
		}
	}

	sort.Ints(res)
	return res
}

// @lc code=end

