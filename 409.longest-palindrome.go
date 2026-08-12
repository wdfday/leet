/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}

	res := 0
	odd := false
	for _, v := range cnt {
		if v%2 == 0 {
			res += v
		} else {
			res += v - 1
			odd = true
		}
	}
	if odd {
		res++
	}
	return res

}

// @lc code=end

