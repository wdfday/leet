/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {

	if len(s) == 0 {
		return t[0]
	}
	wordS := []byte(s)
	wordT := []byte(t)

	slices.Sort(wordS)
	slices.Sort(wordT)

	i := 0
	for ; i < len(wordS); i++ {
		if wordS[i] != wordT[i] {
			return wordT[i]
		}
	}

	return wordT[i]
}

// @lc code=end

