/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

import "strconv"

// @lc code=start
func countAndSay(n int) string {
	if n < 1 {
		return ""
	}

	result := "1"
	for i := 2; i <= n; i++ {
		result = RLE(result)
	}

	return result
}

func RLE(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := make([]byte, 0, len(s)*2)
	count := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
			continue
		}

		result = append(result, strconv.Itoa(count)...)
		result = append(result, s[i-1])
		count = 1
	}

	result = append(result, strconv.Itoa(count)...)
	result = append(result, s[len(s)-1])

	return string(result)

}

// @lc code=end

