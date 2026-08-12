/*
 * @lc app=leetcode id=594 lang=golang
 *
 * [594] Longest Harmonious Subsequence
 */

// @lc code=start
func findLHS(nums []int) int {
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	res := 0
	for v, c := range count {
		if c2, ok := count[v+1]; ok {
			if c+c2 > res {
				res = c + c2
			}
		}
	}

	return res
}

// @lc code=end

