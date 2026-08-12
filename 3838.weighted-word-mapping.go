/*
 * @lc app=leetcode id=3838 lang=golang
 *
 * [3838] Weighted Word Mapping
 */

// @lc code=start
func mapWordWeights(words []string, weights []int) string {
	res := []byte{}
	for _, word := range words {
		sum := 0
		for _, c := range word {
			sum += weights[c-'a']
		}
		res = append(res, byte('z'-sum%26))
	}
	return string(res)

}

// @lc code=end

