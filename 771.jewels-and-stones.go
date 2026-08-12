/*
 * @lc app=leetcode id=771 lang=golang
 *
 * [771] Jewels and Stones
 */

// @lc code=start
func numJewelsInStones(jewels string, stones string) int {
	res := 0
	for _, v := range stones {
		if strings.ContainsRune(jewels, v) {
			res++
		}
	}
	return res

}

// @lc code=end

