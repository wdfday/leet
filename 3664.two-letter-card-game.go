/*
 * @lc app=leetcode id=3664 lang=golang
 *
 * [3664] Two-Letter Card Game
 */

// @lc code=start
func score(cards []string, x byte) int {
	left := make([]int, 10)
	right := make([]int, 10)
	both := 0

	for _, c := range cards {
		if c[0] == x && c[1] == x {
			both++
		} else if c[0] == x {
			left[c[1]-'a']++
		} else if c[1] == x {
			right[c[0]-'a']++
		}
	}

	solve := func(cnt []int, have int) int {
		total := have
		mx := have
		for _, v := range cnt {
			total += v
			if v > mx {
				mx = v
			}
		}
		return min(total/2, total-mx)
	}

	ans := 0
	for i := 0; i <= both; i++ {
		ans = max(ans, solve(left, i)+solve(right, both-i))
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

