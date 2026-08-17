/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}

	code := func(c byte) int {
		switch c {
		case 'A':
			return 0
		case 'C':
			return 1
		case 'G':
			return 2
		case 'T':
			return 3
		}
		return 0
	}

	seen := make(map[int]int)
	ans := []string{}

	window := 0

	for i := 0; i < len(s); i++ {
		window = (window << 2) | code(s[i])

		window &= (1 << 20) - 1

		if i < 9 {
			continue
		}

		seen[window]++

		if seen[window] == 2 {
			ans = append(ans, s[i-9:i+1])
		}
	}

	return ans
}
// @lc code=end

