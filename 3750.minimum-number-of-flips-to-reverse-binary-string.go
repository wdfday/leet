/*
 * @lc app=leetcode id=3750 lang=golang
 *
 * [3750] Minimum Number of Flips to Reverse Binary String
 */

// @lc code=start
func minimumFlips(n int) int {
	res := 0
	s := strconv.FormatInt(int64(n), 2)
    for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			res+=2
		}
    }
	return res
}	
// @lc code=end

