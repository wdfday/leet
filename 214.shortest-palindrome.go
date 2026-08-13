/*
 * @lc app=leetcode id=214 lang=golang
 *
 * [214] Shortest Palindrome
 */

// @lc code=start
func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	// build transformed string: "^#a#b#c#$"
	t := "^"
	for _, c := range s {
		t += "#" + string(c)
	}
	t += "#$"

	n := len(t)
	p := make([]int, n) // p[i] = radius of palindrome centered at i (in t)
	center, right := 0, 0

	for i := 1; i < n-1; i++ {
		mirror := 2*center - i
		if i < right {
			p[i] = min(right-i, p[mirror])
		}
		// try expand
		for t[i+p[i]+1] == t[i-p[i]-1] {
			p[i]++
		}
		// update center/right if expanded past current boundary
		if i+p[i] > right {
			center, right = i, i+p[i]
		}
	}

	// find max k such that palindrome starting at index 0 of s exists
	// i.e. find center i where palindrome touches the left edge (p[i] == i - 1, since t[0]='^')
	maxLen := 0
	for i := 1; i < n-1; i++ {
		if i-p[i] == 1 { // palindrome reaches all the way to start
			if p[i] > maxLen {
				maxLen = p[i]
			}
		}
	}

	// maxLen is length of longest palindromic prefix of s
	suffix := s[maxLen:]
	// reverse suffix
	rev := []byte(suffix)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return string(rev) + s
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	rev := []byte(s)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}

	combined := s + "#" + string(rev)
	n := len(combined)
	fail := make([]int, n)

	for i := 1; i < n; i++ {
		j := fail[i-1]
		for j > 0 && combined[i] != combined[j] {
			j = fail[j-1]
		}
		if combined[i] == combined[j] {
			j++
		}
		fail[i] = j
	}

	k := fail[n-1] // length of longest palindromic prefix of s
	suffix := s[k:]
	revSuffix := make([]byte, len(suffix))
	for i, j := 0, len(suffix)-1; j >= 0; i, j = i+1, j-1 {
		revSuffix[i] = suffix[j]
	}

	return string(revSuffix) + s
}