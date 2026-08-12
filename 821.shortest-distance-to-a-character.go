/*
 * @lc app=leetcode id=821 lang=golang
 *
 * [821] Shortest Distance to a Character
 */

// @lc code=start

func shortestToChar(s string, c byte) []int {
	n := len(s)
	res := make([]int, n)

	for i := range res {
		res[i] = n
	}

	prev := -n
	for i := 0; i < n; i++ {
		if s[i] == c {
			prev = i
		}
		res[i] = i - prev
	}

	prev = 2 * n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			prev = i
		}
		res[i] = min(res[i], prev-i)
	}

	return res
}

// @lc code=end

