/*
 * @lc app=leetcode id=316 lang=golang
 *
 * [316] Remove Duplicate Letters
 */

// @lc code=start
func removeDuplicateLetters(s string) string {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}

	stack := []byte{}
	seen := [26]bool{}

	for _, c := range []byte(s) {
		cnt[c-'a']--

		if seen[c-'a'] {
			continue
		}

		for len(stack) > 0 &&
			stack[len(stack)-1] > c &&
			cnt[stack[len(stack)-1]-'a'] > 0 {

			seen[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, c)
		seen[c-'a'] = true
	}

	return string(stack)
}
// @lc code=end

