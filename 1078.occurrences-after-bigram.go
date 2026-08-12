/*
 * @lc app=leetcode id=1078 lang=golang
 *
 * [1078] Occurrences After Bigram
 */

// @lc code=start
func findOcurrences(text string, first string, second string) []string {
    words := strings.Split(text, " ")
	res := make([]string, 0)
	for i := 0; i < len(words) - 2; i++ {
		if words[i] == first && words[i+1] == second {
			res = append(res, words[i+2])
		}
	}
	return res
}
// @lc code=end

